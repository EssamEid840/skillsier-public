import { z } from "zod";

export const ProfileBasics = z.object({
  fullName: z.string().min(2).max(120),
  headline: z.string().max(120).optional().default(""),
  bio: z.string().max(2000).optional().default(""),
  location: z.string().max(120).optional(),
});

export const ContactEmail = z.object({
  email: z.string().email(),
  primary: z.boolean().default(false),
});

export const Visibility = z.object({
  searchable: z.boolean().default(true),
  profilePublic: z.boolean().default(true)
});

export type ProfileBasicsInput = z.infer<typeof ProfileBasics>;
