Merge Two Folder Structures into a Single fe/ Tree (follow combined-folder-strucure.md style)
=============================================================================================

> Use the **exact style and multi-line comment formatting** shown in combined-folder-strucure.md.**Do not** collapse comments into one line; keep bullets and BE/API lines on their own lines beneath the item.

Input
-----

*   You will receive **two folder trees as attachments** in this conversation.Treat the first as **\[TREE\_A\]** and the second as **\[TREE\_B\]**.
    
*   If the trees are also provided inline between the tags below, use them as the source as well (attachments take precedence if both exist):
    

Plain textANTLR4BashCC#CSSCoffeeScriptCMakeDartDjangoDockerEJSErlangGitGoGraphQLGroovyHTMLJavaJavaScriptJSONJSXKotlinLaTeXLessLuaMakefileMarkdownMATLABMarkupObjective-CPerlPHPPowerShell.propertiesProtocol BuffersPythonRRubySass (Sass)Sass (Scss)SchemeSQLShellSwiftSVGTSXTypeScriptWebAssemblyYAMLXML`   [TREE_A]  (paste the first folder tree here)  [/TREE_A]  [TREE_B]  (paste the second folder tree here)  [/TREE_B]   `

Output Format (Strict)
----------------------

*   Return **one** Markdown **code block** only.
    
*   The code block must contain **one continuous** tree starting at **fe/**.
    
*   Use **box-drawing characters**: │, ├──, └──.
    
*   Inside each folder, list **folders first**, then **files** (**alphabetical**, case-insensitive).
    
*   Keep **multi-line comments** exactly like combined-folder-strucure.md:
    
    *   Optional short title comment on the **same line** after two spaces and #.
        
    *   Additional comment lines appear on **indented subsequent lines**, each starting with #.
        
    *   Use # - ... bullets for features/notes.
        
    *   Keep # BE: ... and each # GET/POST/... endpoint on **separate lines**.
        
*   **No prose** before or after the code block—**output the tree only**.
    

Structure Style (Very Important)
--------------------------------

*   Show a **fully expanded** nested tree (parents with their children beneath them).
    
*   **Do not** write a path as a standalone header (avoid lines like apps/web/src/app/\[locale\]/(dashboard)/ alone).
    
*   Always represent hierarchy using the tree.
    

Placement Rules (Not limited to web/mobile)
-------------------------------------------

Everything must live under the single root **fe/**. Use these canonical areas:

*   **Applications**
    
    *   fe/apps/web/ — Web app (e.g., Next.js).
        
    *   fe/apps/mobile/ — Mobile app (e.g., Expo/React Native).
        
*   **Shared Libraries**
    
    *   fe/packages/ — Shared packages/libs (normalize inputs like pkg/, libs/, shared/ into **packages/**).
        
*   **Documentation**
    
    *   fe/docs/ — Product and technical docs.
        
*   **Automation & Tooling**
    
    *   fe/scripts/ — Scripts/CLIs (build, dev, test, codegen).
        
*   **Configuration**
    
    *   fe/config/ — Centralized configs (env templates, linting bases, etc.).
        
*   **CI/CD & Hooks**
    
    *   fe/.github/ — Workflows, issue/pr templates.
        
    *   fe/.husky/ — Git hooks.
        
    *   fe/.vscode/ — Editor settings.
        
*   **Optional if present**
    
    *   fe/infra/ (IaC/devops), fe/tools/, fe/examples/, fe/tests/, fe/assets/.
        

> When inputs use synonymous top-level names (e.g., pkg/ vs packages/), **normalize to the canonical path above** and **preserve all original comments** under the moved items.

Deduplication & Merge Rules
---------------------------

1.  **Same path duplicates (file or folder):**
    
    *   Show **one** instance **only**.
        
    *   **Merge children** (for folders).
        
    *   **Merge comment blocks** from both sources under the single item:
        
        *   Keep **multi-line style** (do **not** join into one line).
            
        *   **Preserve original order** of comments; remove **exact duplicate lines**.
            
2.  **Platform variants:** If two files share a path but differ by platform suffix (e.g., .web.tsx vs .native.tsx), **keep both**.
    
3.  **File vs Folder collision (same name used as both):**
    
    *   Treat the path as a **folder**.
        
    *   Include the file **inside that folder**, and add on the file’s line:# conflict: file also existed.
        

Comment Preservation & Ordering
-------------------------------

*   Keep wording **exactly** as provided—**no paraphrasing**.
    
*   Preferred order within a comment block:
    
    1.  Short title comment (if present) on the item line
        
    2.  Bullet lines # - ...
        
    3.  \# BE: ... line(s)
        
    4.  HTTP endpoint lines (# GET ..., # POST ..., etc.)
        
*   Maintain indentation so all comment lines visually **belong** to their item.
    

Final Requirement
-----------------

Return **only** the final merged tree inside **one** Markdown code block. **Nothing else.**