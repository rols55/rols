

import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';

export async function CheckSession() {
  const backendEndpoint = `http://${process.env.API_HOST_SERVER_SIDE}:${process.env.API_PORT_SERVER_SIDE}/api/check-session`;
  const cookieData = JSON.stringify({ session_token: cookies().get('session_token')?.value});
  if (!cookieData) {
    return false;
  }
  try {
    const response = await fetch(backendEndpoint, {
      method: "POST",
      body:  cookieData,
    });
    if (response.ok) {
      return true;
    }
  } catch (error) {
    console.log("Check Session Error: "+error);
  }
  return false;
}

export async function middleware(request) {
  const hasSession = await CheckSession();
  const nextUrlPathname = request.nextUrl.pathname;

  if (hasSession && nextUrlPathname.startsWith('/login') || hasSession && nextUrlPathname.startsWith('/register') ) {
    return NextResponse.redirect(new URL('/', request.url));
  }

  if (!hasSession && !nextUrlPathname.startsWith('/login') && !nextUrlPathname.startsWith('/register')) {
    return Response.redirect(new URL('/login', request.url));
  }
}


export const config = {
  matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
}
