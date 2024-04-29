import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

// This function can be marked `async` if using `await` inside
export function middleware(request: NextRequest) {
    const url = request.nextUrl;
    const path = url.pathname;  // Get the current path from the URL
    const token = request.cookies.get('access_token');
    const role = request.cookies.get('role');

    // Define role-based access rules for each path prefix
    const accessControl = {
        '/dashboard': ['Admin'],  // All subpaths in dashboard require Admin role
        '/dashboard/user': ['User', 'Admin'],  // Specific override for this path
    };

    if (!token || !role) {
        // No token or role found, redirect to login
        return NextResponse.redirect(new URL('/auth', request.url));
    }

    // Function to check access based on path prefixes
    const hasAccess = (path: string) => {
        console.log('Checking access for path:', path);

        for (const basePath in accessControl) {
            if (path.startsWith(basePath)) {
                const allowedRoles = accessControl[basePath as keyof typeof accessControl];
                console.log('Allowed roles:', allowedRoles);
                console.log('User role:', role.value);
                console.log(allowedRoles.includes(role.value));
                return allowedRoles.includes( role.value);
            }
        }
        return false;  // Default deny if no matching path prefix found
    };

    if (hasAccess(path)) {

        return NextResponse.next();  // User has access
    } else {
        // User does not have access to this path
        return NextResponse.redirect(new URL('/auth', request.url));
    }
}

export const config = {
    matcher: ['/dashboard/:path*', '/cart/:path*']  // Ensures this middleware runs for these paths
};