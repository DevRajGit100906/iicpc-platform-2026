# IICPC Distributed Benchmarking & Hosting Platform

Contestants upload a **matching engine**; the platform builds it, hosts it in a hardened
sandbox, bombards it with a fleet of bots, validates its fills against a reference oracle,
and ranks it on a **gated composite score** shown on a live leaderboard.

This README is the **file manifest** (what every file is and where it goes), the **run
guide**, and the **completion roadmap**.

---

## Architecture (local / brief-aligned)

```bash
contestant ──upload(zip)──▶ Orchestrator (:8082)
                              │ docker build
                              │ run engine in hardened sandbox (published port, readiness gate)
                              ├── Pass 1: CORRECTNESS  (sequenced)  → correctness.Check vs oracle
                              ├── Pass 2: LOAD          (concurrent) → botfleet, captures latencies
                              │                                        └─telemetry─▶ Kafka ─▶ ingester (live view)
                              ├── scoring.Compute  (gated composite)
                              └── Redis: ZSET live_leaderboard + hash run:<id>
                                            │
                              Leaderboard (:8081) ──WS──▶ React board (:3000)
```

Key point: the **untrusted engine never touches Kafka** — only the platform's own bots emit
telemetry, so the broker networking is trivial and the run-id is authoritative.

---

## File manifest

Legend: **[NEW]** add from this session · **[EDIT]** replace your existing file ·
**[DONE]** already in your repo, keep · **[TODO]** not built yet (see roadmap).

```bash
iicpc-platform-2026/
├── go.mod, go.sum                         [EDIT]  set module path; `go mod tidy`
├── Makefile                               [NEW]   every workflow target
├── docker-compose.yml                     [DONE]  kafka + redis (fine as-is once flipped)
├── README.md                              [NEW]   this file
│
├── pkg/orderbook/
│   ├── orderbook.go                       [NEW]   price-time-priority CLOB — engine core AND oracle
│   └── orderbook_test.go                  [NEW]   8 invariant tests (FIFO, price priority, FOK, IOC, cancel)
│
├── sut/
│   ├── reference-engine/
│   │   ├── main.go                        [NEW]   correct sample engine; reports fills (replaces dummy)
│   │   └── Dockerfile                     [EDIT]  build reference-engine; EXPOSE 8080 (snippet below)
│   └── samples/                           [TODO]  slow+correct, fast+incorrect, malicious (demo "win move")
│
├── services/
│   ├── orchestrator/main.go               [EDIT]  unified: build→host→correctness+load→gated score→Redis
│   ├── bot-fleet/
│   │   ├── botfleet.go                    [NEW]   platform-owned concurrent load generator
│   │   ├── botfleet_test.go               [NEW]   run-id authority + unique seq under -race
│   │   └── sink_kafka.go                  [NEW]   telemetry → Kafka
│   ├── scoring/
│   │   ├── scoring.go                     [NEW]   gated composite + percentiles
│   │   ├── scoring_test.go                [NEW]   gate keeps correct > fast-but-wrong
│   │   └── correctness/
│   │       ├── correctness.go             [NEW]   sequenced diff vs oracle (D.5)
│   │       └── correctness_test.go        [NEW]   catches no-fill + wrong-price engines, localized
│   ├── leaderboard/main.go                [EDIT]  read ranks + per-run breakdown → WS
│   └── ingester/main.go                   [DONE]  optional live telemetry view (no longer the scorer)
│
├── frontend/index.html                    [EDIT]  correctness badge + p99 + TPS + gated score
│
├── proto/iicpc/v1/{messages,services}.proto  [DONE]  contracts — replace YOUR_USERNAME, then `make proto`
├── schemas/iicpc/v1/*.pb.go               [DONE]  generated (regen after fixing go_package)
│
├── docs/adr/ADR-0001-architecture.md      [TODO]  write (blueprint, deliverable #2)
├── docs/adr/ADR-0002-isolation.md         [TODO]  sandbox-vs-latency trade-off
│
└── infra/
    ├── kubernetes/                        [TODO]  manifests / Helm
    └── terraform/                         [TODO]  cluster + node pools (deliverable #3)
```

---

## Setup: module path

Every `[NEW]`/`[EDIT]` `.go` file imports siblings via a placeholder `MODPATH`. Set it once:

```bash
# from the repo root — replace with your real module (check the `module` line in go.mod)
MODULE=$(head -1 go.mod | awk '{print $2}')
grep -rl MODPATH --include=*.go . | xargs sed -i "s#MODPATH#${MODULE}#g"
go mod tidy
```

Also fix the literal `YOUR_USERNAME` in `proto/.../*.proto` and your generated `*.pb.go`
(set `go_package` to your real module), then `make proto`.

`Dockerfile` for `sut/reference-engine` (multi-stage):

```dockerfile
FROM golang:1.22 AS build
WORKDIR /src
COPY . .
RUN go build -o /engine ./sut/reference-engine
FROM gcr.io/distroless/base-debian12
COPY --from=build /engine /engine
EXPOSE 8080
ENTRYPOINT ["/engine"]
```

---

## Run it locally

```bash
make up                 # start kafka + redis
make test               # confirm orderbook/botfleet/correctness/scoring pass

# terminal A — the leaderboard
make leaderboard
# terminal B — the orchestrator (accepts uploads, scores)
make orchestrator
# browser — the live board
make frontend           # open http://localhost:3000

# submit the reference engine as a sample (zip the build context, POST it)
#   the field name is "submission"
curl -F submission=@submission.zip http://localhost:8082/upload
```

A run flows: `BUILDING → DEPLOYING → CHECKING → RUNNING → FINISHED`, and the board row fills
in with correctness %, p99, TPS, and the gated composite.

---

## Roadmap — how to complete it (mapped to the Master Build Plan)

**Done (the spine, world-class where it counts):**

- A.1/A.2 foundations: repo, contracts (proto + buf + generated stubs).
- The vertical slice: upload → build → sandbox → load → live board.
- **B (partial):** hardened sandbox — cap-drop ALL, no-new-privileges, PID/mem/CPU limits, readiness gate.
- **C (core):** platform-owned, seeded, concurrent, deterministic load generator.
- **D.4/D.5:** reference oracle + sequenced differential correctness, localizes the first violation.
- **E.1:** gated composite score + percentiles; **E.2:** live board with the breakdown.

**Next, cheap, high-value (finish the personal-project version):**

1. **Sample submissions** (`sut/samples/`): slow+correct (add a sleep), fast+incorrect (the +1-tick
   engine), malicious (tries egress). Preloading these makes the board visibly *discriminate* — the
   single most convincing demo artifact (Plan: E.4 "win move").
2. **Blueprint + ADRs** (`docs/`): the Master Build Plan is 90% of the text; distill it into
   `ADR-0001` (architecture) and `ADR-0002` (sandbox-vs-latency trade-off). Deliverable #2.
3. **Warmup phase** (B.5): run an unmeasured warmup before the scored window — kills cold-start
   skew, and it's ~10 lines in the orchestrator.
4. **More invariants** (D.5 / R2): quantity conservation, no-crossed-book, sequence monotonicity —
   each is a small pure check over the oracle output.
5. **Minimal IaC** (`infra/`): even a documented `docker-compose` bring-up + a basic k8s manifest
   gestures at deliverable #3.

**Aspirational (the full production build — real but large):**

- Kata/gVisor/Firecracker isolation + dedicated cores/NUMA (B.3/B.4); Cilium deny + Falco (I).
- Hawkes/archetype microstructure + open-loop CO correction + HdrHistogram (C.3/C.5, D.2).
- Linearizability checking with Porcupine (D.5 "elite move").
- Resilience/chaos (F); host/kernel tuning + KEDA autoscale to 10^5 bots (G); Terraform + Helm +
  ArgoCD GitOps (H).

These are genuine systems-engineering projects in their own right — pursue them one epic at a time;
the spine above already works end-to-end so you're never stranded.
