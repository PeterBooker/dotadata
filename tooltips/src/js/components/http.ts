export module HTTP {

    export function Get(url: string, callback: any) {
        var req = new XMLHttpRequest()
        req.responseType = 'json'
        req.addEventListener( 'load', callback )
        req.addEventListener( 'error', callback )
        req.addEventListener( 'abort', callback )
        req.open( 'GET', encodeURI( url ) )
        req.send()

        return req
    }

}