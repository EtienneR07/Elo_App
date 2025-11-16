// Simple router for Svelte 5 with HTML5 history API

// Router state
let currentPath = $state(window.location.pathname);

// Route type
export interface Route {
  path: string;
  component: any;
}

// Navigate function
export function navigate(path: string, replace = false) {
  if (replace) {
    window.history.replaceState({}, '', path);
  } else {
    window.history.pushState({}, '', path);
  }
  currentPath = path;
}

// Initialize router
export function initRouter() {
  // Handle browser back/forward buttons
  window.addEventListener('popstate', () => {
    currentPath = window.location.pathname;
  });

  // Intercept clicks on links
  document.addEventListener('click', (e) => {
    const target = e.target as HTMLElement;
    const anchor = target.closest('a');

    if (anchor && anchor.href && anchor.origin === window.location.origin) {
      const path = anchor.pathname;
      // Only prevent default if it's an internal link
      if (path !== window.location.pathname) {
        e.preventDefault();
        navigate(path);
      }
    }
  });
}

// Get current route
export function getCurrentPath() {
  return currentPath;
}

// Match route
export function matchRoute(routes: Route[]) {
  const path = getCurrentPath();

  for (const route of routes) {
    if (route.path === path) {
      return route.component;
    }
  }

  // Return first route as fallback (usually home)
  return routes[0]?.component;
}
