import Link from "next/link";

export default function NotFound() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-santan-050">
      <div className="text-center space-y-4">
        <h1 className="text-4xl font-bold text-kuali-950">404</h1>
        <p className="text-kuali-700">Halaman tidak ditemukan.</p>
        <Link href="/today" className="text-rempah-500 text-sm font-medium hover:underline">Kembali ke Today</Link>
      </div>
    </div>
  );
}
