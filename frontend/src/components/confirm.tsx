"use client";

export function ConfirmDialog({
  open,
  title,
  message,
  confirmLabel,
  onConfirm,
  onCancel,
}: {
  open: boolean;
  title: string;
  message: string;
  confirmLabel?: string;
  onConfirm: () => void;
  onCancel: () => void;
}) {
  if (!open) return null;

  return (
    <div className="fixed inset-0 bg-ink-950/50 flex items-center justify-center p-4 z-50" onClick={onCancel}>
      <div className="bg-white-000 rounded-xl max-w-sm w-full p-6 space-y-4" onClick={e => e.stopPropagation()} role="alertdialog" aria-label={title}>
        <h3 className="font-semibold text-ink-900">{title}</h3>
        <p className="text-sm text-ink-700">{message}</p>
        <div className="flex gap-2 justify-end">
          <button onClick={onCancel} className="px-4 py-2 text-sm text-ink-700 hover:bg-steel-200 rounded-lg transition-colors">Cancel</button>
          <button onClick={onConfirm} className="px-4 py-2 text-sm bg-feedback-error text-white-000 rounded-lg font-medium hover:opacity-90 transition-colors">{confirmLabel || "Delete"}</button>
        </div>
      </div>
    </div>
  );
}
