# PRD: Favor Tracker v1


| Field        | Value                           |
| ------------ | ------------------------------- |
| Client       | AlienDev                        |
| Product      | Favor Tracker (`favor-tracker`) |
| Status       | Ready for implementation        |
| Owner        | Michael "AlienDev" Youngblood   |
| Last updated | 2026-09-24                      |


## Change history


| Date       | Author           | Change                                                                |
| ---------- | ---------------- | --------------------------------------------------------------------- |
| 2026-09-24 | AlienDev / agent | locked PRD — Ready for implementation (path/dates/hosting remain TBD) |
| 2026-09-24 | AlienDev / agent | Scaffolded PRD; requirements interview started                        |


## Overview

Favor Tracker is a **public Go portfolio project** (source viewable; all rights reserved — not open-source licensed). The demo domain is personal Favor-driving tracking (the owner uses it for real), but the primary product goal is shipping clean, readable public Go — not building a multi-driver SaaS first.

**v1 interface:** CLI with `help` and interactive prompted flows for full trip or full expense entry.

**v1 reporting (three modes, same data):**

1. **CLI** — quick terminal list/summary.
2. **HTML** — interactive local browse in the browser.
3. **CSV** — file export for spreadsheets.

**Shared report capabilities (all modes):**

- Date-range filter: day / week / month / custom
- Trip list: each delivery with earnings, hours, miles
- Expense list: by category
- Totals: gross pay, tips, expenses, net
- Derived stats: per-mile and per-hour
- Breakdown by expense category

No TUI or general-purpose localhost CRUD web app. No cloud sync. No shift entity.

**v1 data model:**

- **Trips** (each Favor delivery): earnings (pay, tips, bonuses, fees), hours/time, miles.
- **Expenses** (first-class): amount + date + fixed category; optional note.

**v1 expense categories (fixed):** Gas/fuel · Maintenance/repairs · Car wash · Parking · Tolls · Supplies · Phone/data · Insurance (portion) · Other (catch-all).

**v1 storage:** single local SQLite database on disk.

**v1 license:** source available / viewable; all rights reserved.

## Success metrics


| Metric               | Baseline | Target                                                  | Notes       |
| -------------------- | -------- | ------------------------------------------------------- | ----------- |
| Help discoverability | None     | `help` documents all commands                           | Round 7     |
| Frictionless entry   | None     | Interactive full trip **or** expense entry              | Round 7     |
| Expense coverage     | None     | Fixed categories in expense flow                        | Round 8     |
| Quick CLI insight    | None     | List/summarize with date range + totals + derived stats | Rounds 9–9b |
| Interactive HTML     | None     | Browse same report contents in HTML                     | Rounds 9–9b |
| CSV export           | None     | Export same report contents to CSV                      | Rounds 9–9b |




## Messaging

Frame as: a small, well-built Go CLI (source available, ARR) that interactively logs Favor trips and categorized expenses into local SQLite, with terminal summaries, interactive HTML browsing, and CSV export — including pay totals, net after expenses, and $/mi / $/hr.

## Timeline / release planning


| Milestone                  | Target date | Notes                      |
| -------------------------- | ----------- | -------------------------- |
| Requirements locked        | TBD         | After interview            |
| MVP usable for own driving | TBD         | Entry + three report modes |
| Public README / release    | TBD         |                            |




## Personas


| Persona                            | Primary?           | Needs / pains                                                |
| ---------------------------------- | ------------------ | ------------------------------------------------------------ |
| Portfolio reader / hiring engineer | Yes (product goal) | Judge Go craft: structure, tests, docs, packaging            |
| Solo Favor driver (owner)          | Yes (demo user)    | Guided entry; CLI / HTML / CSV with totals and derived stats |




## User scenarios



### Scenario 1 — Interactive trip / expense entry

Driver uses prompted CLI flows; data saved to SQLite.

### Scenario 2 — Quick CLI look

Driver filters by day/week/month/custom; sees trip list, expense list, totals (gross, tips, expenses, net), $/mi, $/hr, expense breakdown.

### Scenario 3 — Interactive HTML browse

Same contents as Scenario 2, browsable in local HTML.

### Scenario 4 — CSV export

Same contents exported to CSV for spreadsheets.

### Scenario 5 — Portfolio review

Viewer reads README/LICENSE (ARR), runs help + sample data through CLI/HTML/CSV paths.

## User stories / features / requirements


| Priority | Story / requirement                                                                               | Why it matters  | Notes      |
| -------- | ------------------------------------------------------------------------------------------------- | --------------- | ---------- |
| P0       | As a driver, I want `help`                                                                        | Discoverability | Round 7    |
| P0       | As a driver, I want interactive full **trip** entry                                               | Success bar     | Round 7    |
| P0       | As a driver, I want interactive full **expense** entry with fixed categories                      | Success bar     | Rounds 7–8 |
| P0       | As a driver, I want trips and expenses in local SQLite                                            | Reliability     |            |
| P0       | As a driver, I want **CLI** reports with date range, lists, totals, $/mi, $/hr, expense breakdown | Round 9–9b      |            |
| P0       | As a driver, I want **interactive HTML** with the same report contents                            | Round 9–9b      |            |
| P0       | As a driver, I want **CSV export** with the same report contents                                  | Round 9–9b      |            |
| P1       | As a maintainer, I want README + LICENSE clarifying ARR                                           | Legal clarity   | Round 6    |
| P2       | Non-interactive flags for scripting entry                                                         | Nice-to-have    |            |




## Features out


| Item                                        | Why out                             |
| ------------------------------------------- | ----------------------------------- |
| Multi-driver SaaS / accounts                | Portfolio + solo demo               |
| Full TUI                                    | Deferred; CLI text views instead    |
| General localhost web app (CRUD in browser) | HTML is report/browse, not full app |
| Cloud sync                                  | Deferred                            |
| Shift-level entity                          | Trip + expense only                 |
| Open-source license                         | ARR / source available              |
| Free-form expense categories                | Fixed list for v1                   |
| Zones / vehicle metadata / ratings          | Not selected for MVP                |




## Designs

- Interactive prompt order for trip vs expense: TBD
- HTML interactive browse UX: TBD (tables + filters matching report capabilities)
- CLI column layout: TBD
- CSV schema (one file vs trips.csv + expenses.csv + summary): TBD



## Open issues


| Issue                                                         | Owner    | Status |
| ------------------------------------------------------------- | -------- | ------ |
| How HTML is opened (write file vs `open` vs tiny local serve) | AlienDev | Open   |
| CSV packaging (single vs multiple files)                      | AlienDev | Open   |
| Timeline / dates                                              | AlienDev | Open   |
| Default SQLite path / config                                  | AlienDev | Open   |
| Public hosting (GitHub visibility)                            | AlienDev | Open   |
| Portfolio extras (tests, fixtures, sample DB)                 | AlienDev | Open   |




## Q&A

**Q:** License?
**A:** Source available / viewable; all rights reserved.

**Q:** What does v1 track?
**A:** Trip-level deliveries + categorized expenses in local SQLite.

**Q:** How do I enter data?
**A:** Interactive CLI prompts for full trip or full expense; plus `help`.

**Q:** How do I look at data?
**A:** CLI quick view, interactive HTML browse, and CSV export — all with date ranges, trip/expense lists, totals (gross/tips/expenses/net), $/mi, $/hr, and expense-category breakdown.

**Q:** Expense categories?
**A:** Gas/fuel, maintenance/repairs, car wash, parking, tolls, supplies, phone/data, insurance (portion), other.

**Q:** Shifts?
**A:** No shift entity; aggregate by date range.

## Other considerations

- Code: `favor-tracking` Go module (Hello World today).
- Shared report-query layer feeding CLI, HTML, and CSV is a strong portfolio architecture story.

