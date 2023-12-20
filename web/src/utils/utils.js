
export const GetQuery = function() {
    const url = location.search;
    const query = new Object();
    if (url.indexOf("?") != -1) {
        const str = url.substr(1);
        const strs = str.split("&");
        for(var i = 0; i < strs.length; i ++) {
            query[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
        }
    }
    return query;
}