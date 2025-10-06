import { Card } from "@/components/ui/card";

export default function EmptyState({
  title,
  description,
  action,
  icon,
}: {
  title: string;
  description?: string;
  action?: React.ReactNode;
  icon?: React.ReactNode;
}) {
  return (
    <Card className="p-8 grid place-items-center text-center gap-3">
      {icon ? <div className="text-muted-foreground">{icon}</div> : null}
      <h3 className="text-lg font-medium">{title}</h3>
      {description ? (
        <p className="text-sm text-muted-foreground max-w-prose">{description}</p>
      ) : null}
      {action ?? null}
    </Card>
  );
}
