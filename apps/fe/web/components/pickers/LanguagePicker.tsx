"use client";

import * as React from "react";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import { Plus, X } from "lucide-react";

const LANGS = [
  { code: "en", name: "English" },
  { code: "ar", name: "Arabic" },
  { code: "fr", name: "French" },
  { code: "de", name: "German" },
];

export default function LanguagePicker({
  values,
  onChange,
  max = 10,
}: {
  values: string[];
  onChange: (vals: string[]) => void;
  max?: number;
}) {
  const [open, setOpen] = React.useState(false);

  const remove = (code: string) => onChange(values.filter((v) => v !== code));
  const add = (code: string) => {
    if (!values.includes(code) && values.length < max) onChange([...values, code]);
  };

  return (
    <div className="space-y-2">
      <div className="flex flex-wrap gap-2">
        {values.map((code) => {
          const lang = LANGS.find((l) => l.code === code);
            return (
              <Badge key={code} variant="secondary" className="flex items-center gap-1">
                {lang?.name ?? code}
                <button
                  type="button"
                  onClick={() => remove(code)}
                  className="inline-flex"
                  aria-label={`Remove ${lang?.name ?? code}`}
                >
                  <X className="h-3.5 w-3.5" />
                </button>
              </Badge>
            );
        })}
        {!values.length && <span className="text-sm text-muted-foreground">No languages selected.</span>}
      </div>

      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button type="button" variant="outline" size="sm" className="gap-2">
            <Plus className="h-4 w-4" /> Add language
          </Button>
        </PopoverTrigger>
        <PopoverContent className="p-0 w-[--radix-popover-trigger-width]">
          <Command>
            <CommandInput placeholder="Search language…" />
            <CommandList>
              <CommandEmpty>No results.</CommandEmpty>
              <CommandGroup>
                {LANGS.map((l) => (
                  <CommandItem
                    key={l.code}
                    value={l.name}
                    onSelect={() => {
                      add(l.code);
                      setOpen(false);
                    }}
                  >
                    {l.name}
                  </CommandItem>
                ))}
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
    </div>
  );
}
