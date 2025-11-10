import { MetadataRoute } from 'next';

export default function robots(): MetadataRoute.Robots {
  const baseUrl = 'https://elo-app.com'; // Update with your actual domain

  return {
    rules: [
      {
        userAgent: '*',
        allow: '/',
        disallow: ['/api/', '/admin/'], // Prevent crawling API routes and admin areas
      },
    ],
    sitemap: `${baseUrl}/sitemap.xml`,
  };
}
