Merge Multiple Folder Structures into a Single fe/ Tree (follow combined-folder-strucure.md style)
==================================================================================================

> The **main folder structure** and **styling authority** is combined-folder-strucure.md.Use its exact tree and **multi-line comment** formatting (bullets and BE/API lines). **Do not** collapse comments into one line.

Inputs
------

*   You will receive **multiple** folder-structure documents as attachments in this conversation.
    
    *   Treat **combined-folder-strucure.md** as the **BASE** tree.
        
    *   Treat **all other attached documents** (names/count vary) as **ADDITIVE sources** to be merged into the base.
        
*   If trees are also provided inline, **attachments take precedence**.
    

Analysis (do this **silently**, do **not** print your notes)
------------------------------------------------------------

*   Parse **every** attached document and build a **complete union** of all paths.
    
*   Normalize top-level areas to the canonical layout (see _Canonical Placement_ below).
    
*   For each directory, compute the **sorted** list of subfolders (A–Z) then files (A–Z).
    
*   For duplicates, plan comment merging (multi-line blocks, no paraphrase, dedupe identical lines).
    
*   **Completeness audit:** verify that **every path** from **all** inputs is represented in the final merged tree (except exact duplicates that were deduplicated). If any source path would be dropped by normalization, **relocate** it to the canonical area instead of omitting it.
    

> Do **not** output your analysis/plan. Output only the merged tree (or parts) per the Output Target rules.

Output Target (choose the first method that fully fits)
-------------------------------------------------------

1.  **Preferred:** Write the full merged tree to a single Markdown file named merged-fe-tree.md and return only a link to that file.
    
2.  **Fallback (if file output isn’t possible):** Emit the tree in **multiple consecutive code blocks**, one **top-level folder per block** in this exact order:fe/.github, fe/.husky, fe/.vscode, fe/apps/web, fe/apps/mobile, fe/packages, fe/docs, fe/scripts, fe/config, fe/infra, fe/tools, fe/examples, fe/tests, fe/assets, then **root files**.Prefix each block with an HTML comment header . Continue with subsequent parts **until the entire tree is printed**.
    
3.  **Only if it truly fits in one message:** emit **one** Markdown code block containing the **entire** tree.
    

> **Never truncate.** If you approach length limits, immediately continue with the next PART until finished.

Rendering Rules
---------------

*   Use **box-drawing characters**: │, ├──, └──.
    
*   One unified tree rooted at **fe/**.
    
*   Inside each folder, list **folders first**, then **files** (alphabetical, case-insensitive).
    
*   Keep **multi-line comments** exactly like combined-folder-strucure.md:
    
    *   Optional short title on the **same line** after two spaces and #.
        
    *   Additional comments on **indented lines**, each starting with #.
        
    *   Use # - ... bullets; keep # BE: ... and each # GET/POST/... on their own lines.
        
*   Preserve special route segments (e.g., \[locale\], (dashboard)) **as-is**.
    

Canonical Placement (not limited to web/mobile)
-----------------------------------------------

Everything lives under **fe/**. Normalize top-level areas to:

*   **Applications**
    
    *   fe/apps/web/ — Web app (e.g., Next.js)
        
    *   fe/apps/mobile/ — Mobile app (e.g., Expo/React Native)
        
*   **Shared Libraries**
    
    *   fe/packages/ — Shared packages/libs_(normalize inputs like pkg/, libs/, shared/ to packages/, preserving all comments)_
        
*   **Documentation** — fe/docs/
    
*   **Automation & Tooling** — fe/scripts/
    
*   **Configuration** — fe/config/
    
*   **CI/CD & Hooks** — fe/.github/, fe/.husky/, fe/.vscode/
    
*   **Optional (if present in sources)** — fe/infra/, fe/tools/, fe/examples/, fe/tests/, fe/assets/
    

Merge & Dedup Rules (NO MISSING ITEMS)
--------------------------------------

1.  **Base + Multiple Additives**
    
    *   Start from the base. Merge **all other** provided trees **into** the base.
        
    *   **Do not miss any folder or file.** The **only** items you may omit are **exact duplicates** that are deduplicated.
        
2.  **Same path duplicates (folder or file)**
    
    *   Show **one** instance **only**.
        
    *   **Folders:** merge children from all sources.
        
    *   **Files:** keep **one** file and **merge all comment blocks** in multi-line style (append in the order encountered; remove **exact duplicate** comment lines; never paraphrase).
        
3.  **Platform variants**
    
    *   If files differ by platform suffix (e.g., .web.tsx vs .native.tsx), **keep both**.
        
4.  **File vs Folder name collision** (same path used as both)
    
    *   Treat the path as a **folder**.
        
    *   Place the file **inside** that folder, and add on the file’s line: # conflict: file also existed.
        
5.  **Path normalization**
    
    *   Apply _Canonical Placement_ to relocate synonymous top-level dirs (e.g., pkg/ → packages/) **without dropping anything**.
        

Completeness & Quality Checklist (perform silently before emitting)
-------------------------------------------------------------------

*   Every attached document (except the base) has been fully merged as an additive source.
    
*   Every unique source path appears in the final tree (or in a deduplicated/normalized location).
    
*   All duplicates are merged; no duplicate lines remain.
    
*   Comments are preserved in multi-line style; no paraphrasing; endpoints intact.
    
*   Ordering is alphabetical: **folders first**, then **files**.
    
*   Root is fe/; all canonical areas placed correctly.
    
*   If output length was exceeded, continued with **PART X/Y** blocks until **complete**.
    

Final Requirement
-----------------

Return **only** the merged tree (file link or ordered PARTS as above). **No summaries, no analysis, no extra text.**