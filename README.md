# slp Parser

Module: go.dalton.dog/slp

**This is going to be targeting v3.0.0 support for now**

## Data Processing

### File Parsing: Turning an slp file into initial Go structs

- File: Represents a loaded and initially-parsed file
    - Filepath
    - ParsedRaw
    - ParsedMetadata

- ParsedRaw
    - []Events

- ParsedMetadata
    - StartAt
    - LastFrame
    - Players
    - PlayedOn
    - ConsoleNick
    - Players (map)

### Game-Relevant Parsing

- Game
    - StartInfo
    - EndInfo
    - Stage
    - []Players
    - []Frames
    - Metadata

- StartInfo
    - SlippiVersion
    - IsFrozenPS
    - IsPAL
    - RandomSeed

- EndInfo
    - EndMethod
    - PortLRAS

- Player
    - Character
    - Post
    - Costume
    - StartingStocks
    - Team
    - IsCPU
    - Tag
    - DashbackUCF
    - ShieldDropUCF

- Frame
    - Index
    - []Ports
    - []Items

- Port
    - []PreFrameInfo
    - []PostFrameInfo
