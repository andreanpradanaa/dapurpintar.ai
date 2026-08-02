export function Skeleton({ className = "" }: { className?: string }) {
  return <div className={`animate-pulse bg-bambu-200 rounded ${className}`} />;
}

export function CardSkeleton() {
  return (
    <div className="bg-white border border-bambu-200 rounded-lg p-4 space-y-3">
      <Skeleton className="h-5 w-3/4" />
      <Skeleton className="h-4 w-1/2" />
      <Skeleton className="h-3 w-full" />
    </div>
  );
}

export function ListSkeleton({ count = 3 }: { count?: number }) {
  return (
    <div className="space-y-2">
      {Array.from({ length: count }).map((_, i) => (
        <CardSkeleton key={i} />
      ))}
    </div>
  );
}

export function StatSkeleton() {
  return (
    <div className="bg-white border border-bambu-200 rounded-xl p-4 space-y-2">
      <Skeleton className="h-4 w-1/2" />
      <Skeleton className="h-8 w-1/3" />
    </div>
  );
}
