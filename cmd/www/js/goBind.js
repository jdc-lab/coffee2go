godata = {};

async function go(name, ...params) {
    if (window[name] != undefined) {
        // lorca
        return window[name].apply(this, params);
    }

    // webview
    // if (window[name + "Webview"] != undefined) {
    name = name + "Webview"

    console.log(name)
    console.log(this.toString())
    console.log(this.toString())
    this[name].goCall(params);

    return this[name + "Webview"].data.returning;
    // }
}