import type { Metadata } from "next";
import { Sidebar } from "@/components/sidebar/Sidebar";
import "./globals.css";

export const metadata: Metadata = {
  title: "Palette Agentic",
  description: "Agentic Palette management assistant",
  icons: {
    icon: "/favicon.png",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="h-full antialiased">
      <head>
        <link
          href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap"
          rel="stylesheet"
        />
      </head>
      <body className="min-h-full flex">
        <Sidebar />
        <main className="flex-1 flex flex-col">{children}</main>
      </body>
    </html>
  );
}
