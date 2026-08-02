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
    <div className="fixed inset-0 bg-kuali-950/50 flex items-center justify-center p-4 z-50" onClick={onCancel}>
      <div className="bg-white rounded-xl max-w-sm w-full p-6 space-y-4" onClick={e => e.stopPropagation()} role="alertdialog" aria-label={title}>
        <h3 className="font-semibold text-kuali-950">{title}</h3>
        <p className="text-sm text-kuali-700">{message}</p>
        <div className="flex gap-2 justify-end">
          <button onClick={onCancel} className="px-4 py-2 text-sm text-kuali-700 hover:bg-bambu-200 rounded-lg transition-colors">Cancel</button>
          <button onClick={onConfirm} className="px-4 py-2 text-sm bg-rempah-500 text-white-000 rounded-lg font-medium hover:opacity-90 transition-colors">{confirmLabel || "Delete"}</button>
        </div>
      </div>
    </div>
  );
}
