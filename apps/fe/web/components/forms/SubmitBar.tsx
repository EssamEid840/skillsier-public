"use client";

import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

export default function SubmitBar({
  submitLabel = "Save changes",
  cancelLabel = "Cancel",
  loading = false,
  onCancel,
}: {
  submitLabel?: string;
  cancelLabel?: string;
  loading?: boolean;
  onCancel?: () => void;
}) {
  const router = useRouter();
  return (
    <div className="sticky bottom-0 left-0 right-0 bg-background/80 backdrop-blur supports-[backdrop-filter]:bg-background/60 border-t p-3 flex gap-2 justify-end">
      <Button
        type="button"
        variant="outline"
        onClick={() => (onCancel ? onCancel() : router.back())}
      >
        {cancelLabel}
      </Button>
      <Button type="submit" disabled={loading}>
        {loading ? "Saving…" : submitLabel}
      </Button>
    </div>
  );
}
