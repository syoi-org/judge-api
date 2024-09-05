from aiohttp import web
import asyncio
import nodriver as uc


async def handle(request: web.Request):
    url = request.query.get('url')
    url = url if url else 'https://github.com/syoi-org/judge-api'
    lock: asyncio.Lock = request.app['browser-lock']
    async with lock:
        browser: uc.Browser = request.app['browser']
        page = await browser.get(url)
        return web.Response(text=await page.get_content(), content_type="text/html")

async def browser_context(app: web.Application):
    browser = await uc.start()
    app['browser'] = browser
    app['browser-lock'] = asyncio.Lock()

    yield

    browser.stop()

if __name__ == "__main__":
    app = web.Application()
    app.cleanup_ctx.append(browser_context)
    app.add_routes([web.get("/", handle)])
    web.run_app(app)
