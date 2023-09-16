import { createServer } from 'node:http';
import { createIPX, createIPXMiddleware } from 'ipx';
import { createApp, fromNodeMiddleware, toNodeListener } from 'h3';

const ipx = createIPX({});
const ipxMiddleware = createIPXMiddleware(ipx);

const app = createApp().use('/', fromNodeMiddleware(ipxMiddleware));

createServer(toNodeListener(app)).listen(process.env.PORT || 8080);
