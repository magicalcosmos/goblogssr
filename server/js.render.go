package server

import (
	"strings"
)

const renderJsContent = `
(function() {
const context = RENDER_CONTEXT;
serverBundle(context).then((app) => {
	try {
		const meta = context.meta
		meta.State = JSON.stringify(context.state)

		v8worker.send(83, JSON.stringify(meta), context.v8reqId)

		//console.log("renderToString begin...")
		renderToString(app, context, (err, html) => {
			try {
				if (err) {
					console.error(err)
					if (err.code == 404) {
						v8worker.send(81, "404 Page not found", context.v8reqId)
					} else if (err.code) {
						v8worker.send(81, err.code + " Internal Server Error", context.v8reqId)
					} else {
						v8worker.send(81, err, context.v8reqId)
					}
				} else {
					if (context.styles) {
						v8worker.send(82, context.styles, context.v8reqId)
					}
					v8worker.send(80, html, context.v8reqId)
				}
			} catch(e) {
				console.error(e);
				v8worker.send(81, e, context.v8reqId)
			}
		})
	} catch(e) {
		console.error(e);
		v8worker.send(81, e, context.v8reqId)
	}
}).catch((err) => {
	console.error(err)
	if (err.code == 404) {
		v8worker.send(81, "404 Page not found", context.v8reqId)
	} else {
		v8worker.send(81, err, context.v8reqId)
	}
});
})();
`

var renderJsPart1 string
var renderJsPart2 string
var renderJsLength int

func init() {
	s := "RENDER_CONTEXT"
	idx := strings.Index(renderJsContent, s)
	renderJsPart1 = renderJsContent[:idx]
	renderJsPart2 = renderJsContent[idx+len(s):]
	renderJsLength = len(renderJsPart1) + len(renderJsPart2)
}
