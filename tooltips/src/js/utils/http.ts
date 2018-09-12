export module HTTP {

    export function get(options: string[]) {

        if (!options.url) throw 'Must include URL in GET request.'
        
            options.success = options.success || noop
            options.error = options.error || noop
        
            var xhr = new XMLHttpRequest()
            xhr.open('GET', encodeURI(options.url))
            xhr.onload = function () {
                if (xhr.status === 200) {
                    options.success(xhr.responseText)
                } else {
                    options.error(xhr.responseText)
                }
            };
            xhr.send()

    }

}