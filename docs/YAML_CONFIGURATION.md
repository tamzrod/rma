# YAML CONFIGURATION

## 1. Purpose
- Declarative description of memory topology and listeners
- Zero runtime logic
- Zero policy execution

## 2. Design Principles
- Explicit over implicit
- Absence implies openness
- No speculative fields
- No protocol coupling

## 3. Top-Level Structure
- memory
- listeners

## 4. Memory Banks
- Unique numeric id
- Fixed unit_width_bits
- Fixed units
- Optional init
- Optional access

## 5. init Block
- Declares sealed memory
- Controls who may INIT
- Evaluated only during bootstrap
- CIDR-based declaration only

## 6. access Block
- Restricts runtime access
- CIDR-based declaration only
- No operation semantics inside YAML

## 7. Inheritance Rules
- INIT authority implies ACCESS authority
- ACCESS never implies INIT
- Effective access equals init union access

## 8. Default Behavior
- No init, no access means open for all
- init only means sealed and restricted
- access only means open but restricted

## 9. Listeners
- Raw data-plane endpoints
- No protocol type enumeration
- No authority or lifecycle role

## 10. Explicit Non-Goals
- No auth engine
- No sealing mechanics
- No protocol definitions
- No lifecycle control flags
