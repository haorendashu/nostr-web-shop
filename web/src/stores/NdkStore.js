import { defineStore } from "pinia"
import NDK, { NDKNip07Signer, NDKEvent } from "@nostr-dev-kit/ndk"

export const NdkStore = defineStore({
    id: 'NdkStore',
    state: () => {
        return {
            initialized: false,
            relayUrls: ["wss://nos.lol", "wss://relay.damus.io"],
            ndk: null,
            signer: null,
            profileCache: {},
        }
    },
    actions: {
        async initNdk() {
            if (this.ndk === null) {
                if (window.nostr == null) {
                    // nip07 not support!
                    this.ndk = new NDK({
                        explicitRelayUrls: this.relayUrls,
                        autoConnectUserRelays: false,
                    })
                } else {
                    // nip07 support!
                    this.signer = new NDKNip07Signer();
                    this.ndk = new NDK({
                        signer: this.signer,
                        explicitRelayUrls: this.relayUrls,
                        autoConnectUserRelays: false,
                    })
                }
                await this.ndk.connect()
                this.initialized = true
            }
        },
        async fetchProfile(pubkey) {
            const c = this.profileCache[pubkey];
            if (c) {
                return c;
            }
            const ndkUser = this.ndk.getUser({hexpubkey: pubkey})
            const profile = await ndkUser.fetchProfile()
            if (profile) {
                this.profileCache[pubkey] = profile;
            }
            return profile;
        },
        async pubkey() {
            if (this.ndk.signer) {
                const user = await this.ndk.signer.user();
                if (user) {
                    return user.pubkey
                }
            }

            return "";
        },
        async genNIP98EventText(url, method) {
            const ndkEvent = new NDKEvent(this.ndk)
            ndkEvent.kind = 27235
            ndkEvent.tags = [
                ["u", url],
                ["method", method]
            ]
            await ndkEvent.sign()

            return JSON.stringify(ndkEvent.rawEvent())
        }
    }
})

export function pubkeyShort(pubkey) {
    let length = pubkey.length;
    if (length > 6) {
        return pubkey.substring(0, 5) + ":" + pubkey.substring(length - 5)
    }
    return pubkey
}