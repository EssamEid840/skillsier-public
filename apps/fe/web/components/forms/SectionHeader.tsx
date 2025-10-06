import { Separator } from "@/components/ui/separator";

export default function SectionHeader({
  title,
  description,
}: { title: string; description?: string }) {
  return (
    <div className="space-y-2">
      <div>
        <h2 className="text-xl font-semibold">{title}</h2>
        {description ? (
          <p className="text-sm text-muted-foreground">{description}</p>
        ) : null}
      </div>
      <Separator />
    </div>
  );
}
