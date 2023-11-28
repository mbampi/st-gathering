# Data Gathering Spanning Tree

Opcao E
Algoritmo de coleta usando ST com pais e filhos
- Modificacão: nó não-folha envia ao pai imediatamente cada mensagem recebida
de um filho
  - \# de mensagens?
  - Tempo?
  - depende do uso?


└── Node ID: 0
    ├── Node ID: 1
    │   ├── Node ID: 3
    │   └── Node ID: 4
    └── Node ID: 2
        ├── Node ID: 5
        ├── Node ID: 6
        └── Node ID: 7

### Data Gathering Spanning Tree
Total data gathered by the root (Node 0): 8
Total messages sent: 7
Total time: 145.166µs

### Data Gathering Spanning Tree forwarding data straight to the parent
Total data gathered by the root (Node 0): 8
Total messages sent: 12
Total time: 405.166µs


### Number of Messages
| # Nodes   | Data Gathering ST | Data Gathering ST forwarding data to parent |
| --------- | ----------------- | ------------------------------------------- |
| 3         | 2                 | 2                                           |
| 5         | 4                 | 6                                           |
| 10        | 9                 | 21                                          |
| 20        | 19                | 59                                          |
| 40        | 39                | 155                                         |
| 80        | 79                | 387                                         |
| 160       | 159               | 931                                         |
| 256       | 255               | 1,667                                       |
| 320       | 319               | 2,179                                       |
| 640       | 639               | 4,747                                       |
| 1,280     | 1,279             | 10,764                                      |
| 5,000     | 4,999             | 51,822                                      |
| 10,000    | 9,999             | 113,631                                     |
| 50,000    | 49,999            | 684,481                                     |
| 500,000   | 499,999           | 8,475,732                                   |
| 1,000,000 | 999,999           | 17,951,445                                  | x17.95 |



### Time
| # Nodes   | Data Gathering ST | Data Gathering ST forwarding data to parent |
| --------- | ----------------- | ------------------------------------------- |
| 3         | 119.375µs         | 143.125µs                                   |
| 5         | 131.583µs         | 237.083µs                                   |
| 10        | 270.417µs         | 394.75µs                                    |
| 20        | 326.916µs         | 870.958µs                                   |
| 40        | 694.916µs         | 1.610583ms                                  |
| 80        | 1.3635ms          | 4.060209ms                                  |
| 160       | 2.210042ms        | 14.107625ms                                 |
| 256       | 4.773833ms        | 23.863417ms                                 |
| 320       | 5.11475ms         | 32.046375ms                                 |
| 640       | 13.093708ms       | 56.361333ms                                 |
| 1,280     | 26.571166ms       | 92.647209ms                                 |
| 5,000     | 88.373708ms       | 486.880167ms                                |
| 10,000    | 168.962625ms      | 1.0585515s                                  |
| 50,000    | 808.111917ms      | 6.115064542s                                |
| 500,000   | 8.105962584s      | 1m13.704364917s                             |
| 1,000,000 | 21.150660042s     | 2m47.6733755s                               | x7.9 |



### Trees used

#### 3 nodes

└── Node ID: 0
    ├── Node ID: 1
    └── Node ID: 2

#### 5 nodes

└── Node ID: 0
    ├── Node ID: 1
    │   └── Node ID: 3
    └── Node ID: 2
        └── Node ID: 4

#### 10 nodes

└── Node ID: 0
    ├── Node ID: 1
    │   ├── Node ID: 6
    │   └── Node ID: 3
    │       └── Node ID: 5
    │           ├── Node ID: 8
    │           └── Node ID: 9
    └── Node ID: 2
        ├── Node ID: 7
        └── Node ID: 4

#### 20 nodes

└── Node ID: 0
    ├── Node ID: 1
    │   ├── Node ID: 6
    │   │   ├── Node ID: 10
    │   │   └── Node ID: 11
    │   └── Node ID: 3
    │       └── Node ID: 5
    │           ├── Node ID: 8
    │           │   ├── Node ID: 14
    │           │   └── Node ID: 15
    │           └── Node ID: 9
    │               ├── Node ID: 16
    │               └── Node ID: 17
    └── Node ID: 2
        ├── Node ID: 7
        │   ├── Node ID: 12
        │   └── Node ID: 13
        └── Node ID: 4
            ├── Node ID: 18
            └── Node ID: 19

#### 40 nodes

└── Node ID: 0
    ├── Node ID: 1
    │   ├── Node ID: 6
    │   │   ├── Node ID: 10
    │   │   │   ├── Node ID: 20
    │   │   │   └── Node ID: 21
    │   │   └── Node ID: 11
    │   │       ├── Node ID: 22
    │   │       └── Node ID: 23
    │   └── Node ID: 3
    │       └── Node ID: 5
    │           ├── Node ID: 8
    │           │   ├── Node ID: 14
    │           │   │   ├── Node ID: 28
    │           │   │   └── Node ID: 29
    │           │   └── Node ID: 15
    │           │       ├── Node ID: 30
    │           │       └── Node ID: 31
    │           └── Node ID: 9
    │               ├── Node ID: 16
    │               │   ├── Node ID: 32
    │               │   └── Node ID: 33
    │               └── Node ID: 17
    │                   ├── Node ID: 34
    │                   └── Node ID: 35
    └── Node ID: 2
        ├── Node ID: 7
        │   ├── Node ID: 12
        │   │   ├── Node ID: 24
        │   │   └── Node ID: 25
        │   └── Node ID: 13
        │       ├── Node ID: 26
        │       └── Node ID: 27
        └── Node ID: 4
            ├── Node ID: 18
            │   ├── Node ID: 36
            │   └── Node ID: 37
            └── Node ID: 19
                ├── Node ID: 38
                └── Node ID: 39

#### 80 nodes

└── Node ID: 0
    ├── Node ID: 1
    │   ├── Node ID: 6
    │   │   ├── Node ID: 10
    │   │   │   ├── Node ID: 20
    │   │   │   │   ├── Node ID: 40
    │   │   │   │   └── Node ID: 41
    │   │   │   └── Node ID: 21
    │   │   │       ├── Node ID: 42
    │   │   │       └── Node ID: 43
    │   │   └── Node ID: 11
    │   │       ├── Node ID: 22
    │   │       │   ├── Node ID: 44
    │   │       │   └── Node ID: 45
    │   │       └── Node ID: 23
    │   │           ├── Node ID: 46
    │   │           └── Node ID: 47
    │   └── Node ID: 3
    │       └── Node ID: 5
    │           ├── Node ID: 8
    │           │   ├── Node ID: 14
    │           │   │   ├── Node ID: 28
    │           │   │   │   ├── Node ID: 56
    │           │   │   │   └── Node ID: 57
    │           │   │   └── Node ID: 29
    │           │   │       ├── Node ID: 58
    │           │   │       └── Node ID: 59
    │           │   └── Node ID: 15
    │           │       ├── Node ID: 30
    │           │       │   ├── Node ID: 60
    │           │       │   └── Node ID: 61
    │           │       └── Node ID: 31
    │           │           ├── Node ID: 62
    │           │           └── Node ID: 63
    │           └── Node ID: 9
    │               ├── Node ID: 16
    │               │   ├── Node ID: 32
    │               │   │   ├── Node ID: 64
    │               │   │   └── Node ID: 65
    │               │   └── Node ID: 33
    │               │       ├── Node ID: 66
    │               │       └── Node ID: 67
    │               └── Node ID: 17
    │                   ├── Node ID: 34
    │                   │   ├── Node ID: 68
    │                   │   └── Node ID: 69
    │                   └── Node ID: 35
    │                       ├── Node ID: 70
    │                       └── Node ID: 71
    └── Node ID: 2
        ├── Node ID: 7
        │   ├── Node ID: 12
        │   │   ├── Node ID: 24
        │   │   │   ├── Node ID: 48
        │   │   │   └── Node ID: 49
        │   │   └── Node ID: 25
        │   │       ├── Node ID: 50
        │   │       └── Node ID: 51
        │   └── Node ID: 13
        │       ├── Node ID: 26
        │       │   ├── Node ID: 52
        │       │   └── Node ID: 53
        │       └── Node ID: 27
        │           ├── Node ID: 54
        │           └── Node ID: 55
        └── Node ID: 4
            ├── Node ID: 18
            │   ├── Node ID: 36
            │   │   ├── Node ID: 72
            │   │   └── Node ID: 73
            │   └── Node ID: 37
            │       ├── Node ID: 74
            │       └── Node ID: 75
            └── Node ID: 19
                ├── Node ID: 38
                │   ├── Node ID: 76
                │   └── Node ID: 77
                └── Node ID: 39
                    ├── Node ID: 78
                    └── Node ID: 79

#### 160 nodes

└── Node ID: 0
    ├── Node ID: 1
    │   ├── Node ID: 6
    │   │   ├── Node ID: 10
    │   │   │   ├── Node ID: 20
    │   │   │   │   ├── Node ID: 40
    │   │   │   │   │   ├── Node ID: 80
    │   │   │   │   │   └── Node ID: 81
    │   │   │   │   └── Node ID: 41
    │   │   │   │       ├── Node ID: 82
    │   │   │   │       └── Node ID: 83
    │   │   │   └── Node ID: 21
    │   │   │       ├── Node ID: 42
    │   │   │       │   ├── Node ID: 84
    │   │   │       │   └── Node ID: 85
    │   │   │       └── Node ID: 43
    │   │   │           ├── Node ID: 86
    │   │   │           └── Node ID: 87
    │   │   └── Node ID: 11
    │   │       ├── Node ID: 22
    │   │       │   ├── Node ID: 44
    │   │       │   │   ├── Node ID: 88
    │   │       │   │   └── Node ID: 89
    │   │       │   └── Node ID: 45
    │   │       │       ├── Node ID: 90
    │   │       │       └── Node ID: 91
    │   │       └── Node ID: 23
    │   │           ├── Node ID: 46
    │   │           │   ├── Node ID: 92
    │   │           │   └── Node ID: 93
    │   │           └── Node ID: 47
    │   │               ├── Node ID: 94
    │   │               └── Node ID: 95
    │   └── Node ID: 3
    │       └── Node ID: 5
    │           ├── Node ID: 8
    │           │   ├── Node ID: 14
    │           │   │   ├── Node ID: 28
    │           │   │   │   ├── Node ID: 56
    │           │   │   │   │   ├── Node ID: 112
    │           │   │   │   │   └── Node ID: 113
    │           │   │   │   └── Node ID: 57
    │           │   │   │       ├── Node ID: 114
    │           │   │   │       └── Node ID: 115
    │           │   │   └── Node ID: 29
    │           │   │       ├── Node ID: 58
    │           │   │       │   ├── Node ID: 116
    │           │   │       │   └── Node ID: 117
    │           │   │       └── Node ID: 59
    │           │   │           ├── Node ID: 118
    │           │   │           └── Node ID: 119
    │           │   └── Node ID: 15
    │           │       ├── Node ID: 30
    │           │       │   ├── Node ID: 60
    │           │       │   │   ├── Node ID: 120
    │           │       │   │   └── Node ID: 121
    │           │       │   └── Node ID: 61
    │           │       │       ├── Node ID: 122
    │           │       │       └── Node ID: 123
    │           │       └── Node ID: 31
    │           │           ├── Node ID: 62
    │           │           │   ├── Node ID: 124
    │           │           │   └── Node ID: 125
    │           │           └── Node ID: 63
    │           │               ├── Node ID: 126
    │           │               └── Node ID: 127
    │           └── Node ID: 9
    │               ├── Node ID: 16
    │               │   ├── Node ID: 32
    │               │   │   ├── Node ID: 64
    │               │   │   │   ├── Node ID: 128
    │               │   │   │   └── Node ID: 129
    │               │   │   └── Node ID: 65
    │               │   │       ├── Node ID: 130
    │               │   │       └── Node ID: 131
    │               │   └── Node ID: 33
    │               │       ├── Node ID: 66
    │               │       │   ├── Node ID: 132
    │               │       │   └── Node ID: 133
    │               │       └── Node ID: 67
    │               │           ├── Node ID: 134
    │               │           └── Node ID: 135
    │               └── Node ID: 17
    │                   ├── Node ID: 34
    │                   │   ├── Node ID: 68
    │                   │   │   ├── Node ID: 136
    │                   │   │   └── Node ID: 137
    │                   │   └── Node ID: 69
    │                   │       ├── Node ID: 138
    │                   │       └── Node ID: 139
    │                   └── Node ID: 35
    │                       ├── Node ID: 70
    │                       │   ├── Node ID: 140
    │                       │   └── Node ID: 141
    │                       └── Node ID: 71
    │                           ├── Node ID: 142
    │                           └── Node ID: 143
    └── Node ID: 2
        ├── Node ID: 7
        │   ├── Node ID: 12
        │   │   ├── Node ID: 24
        │   │   │   ├── Node ID: 48
        │   │   │   │   ├── Node ID: 96
        │   │   │   │   └── Node ID: 97
        │   │   │   └── Node ID: 49
        │   │   │       ├── Node ID: 98
        │   │   │       └── Node ID: 99
        │   │   └── Node ID: 25
        │   │       ├── Node ID: 50
        │   │       │   ├── Node ID: 100
        │   │       │   └── Node ID: 101
        │   │       └── Node ID: 51
        │   │           ├── Node ID: 102
        │   │           └── Node ID: 103
        │   └── Node ID: 13
        │       ├── Node ID: 26
        │       │   ├── Node ID: 52
        │       │   │   ├── Node ID: 104
        │       │   │   └── Node ID: 105
        │       │   └── Node ID: 53
        │       │       ├── Node ID: 106
        │       │       └── Node ID: 107
        │       └── Node ID: 27
        │           ├── Node ID: 54
        │           │   ├── Node ID: 108
        │           │   └── Node ID: 109
        │           └── Node ID: 55
        │               ├── Node ID: 110
        │               └── Node ID: 111
        └── Node ID: 4
            ├── Node ID: 18
            │   ├── Node ID: 36
            │   │   ├── Node ID: 72
            │   │   │   ├── Node ID: 144
            │   │   │   └── Node ID: 145
            │   │   └── Node ID: 73
            │   │       ├── Node ID: 146
            │   │       └── Node ID: 147
            │   └── Node ID: 37
            │       ├── Node ID: 74
            │       │   ├── Node ID: 148
            │       │   └── Node ID: 149
            │       └── Node ID: 75
            │           ├── Node ID: 150
            │           └── Node ID: 151
            └── Node ID: 19
                ├── Node ID: 38
                │   ├── Node ID: 76
                │   │   ├── Node ID: 152
                │   │   └── Node ID: 153
                │   └── Node ID: 77
                │       ├── Node ID: 154
                │       └── Node ID: 155
                └── Node ID: 39
                    ├── Node ID: 78
                    │   ├── Node ID: 156
                    │   └── Node ID: 157
                    └── Node ID: 79
                        ├── Node ID: 158
                        └── Node ID: 159

#### 256 nodes

└── Node ID: 0
    ├── Node ID: 1
    │   ├── Node ID: 6
    │   │   ├── Node ID: 10
    │   │   │   ├── Node ID: 20
    │   │   │   │   ├── Node ID: 40
    │   │   │   │   │   ├── Node ID: 80
    │   │   │   │   │   │   ├── Node ID: 160
    │   │   │   │   │   │   └── Node ID: 161
    │   │   │   │   │   └── Node ID: 81
    │   │   │   │   │       ├── Node ID: 162
    │   │   │   │   │       └── Node ID: 163
    │   │   │   │   └── Node ID: 41
    │   │   │   │       ├── Node ID: 82
    │   │   │   │       │   ├── Node ID: 164
    │   │   │   │       │   └── Node ID: 165
    │   │   │   │       └── Node ID: 83
    │   │   │   │           ├── Node ID: 166
    │   │   │   │           └── Node ID: 167
    │   │   │   └── Node ID: 21
    │   │   │       ├── Node ID: 42
    │   │   │       │   ├── Node ID: 84
    │   │   │       │   │   ├── Node ID: 168
    │   │   │       │   │   └── Node ID: 169
    │   │   │       │   └── Node ID: 85
    │   │   │       │       ├── Node ID: 170
    │   │   │       │       └── Node ID: 171
    │   │   │       └── Node ID: 43
    │   │   │           ├── Node ID: 86
    │   │   │           │   ├── Node ID: 172
    │   │   │           │   └── Node ID: 173
    │   │   │           └── Node ID: 87
    │   │   │               ├── Node ID: 174
    │   │   │               └── Node ID: 175
    │   │   └── Node ID: 11
    │   │       ├── Node ID: 22
    │   │       │   ├── Node ID: 44
    │   │       │   │   ├── Node ID: 88
    │   │       │   │   │   ├── Node ID: 176
    │   │       │   │   │   └── Node ID: 177
    │   │       │   │   └── Node ID: 89
    │   │       │   │       ├── Node ID: 178
    │   │       │   │       └── Node ID: 179
    │   │       │   └── Node ID: 45
    │   │       │       ├── Node ID: 90
    │   │       │       │   ├── Node ID: 180
    │   │       │       │   └── Node ID: 181
    │   │       │       └── Node ID: 91
    │   │       │           ├── Node ID: 182
    │   │       │           └── Node ID: 183
    │   │       └── Node ID: 23
    │   │           ├── Node ID: 46
    │   │           │   ├── Node ID: 92
    │   │           │   │   ├── Node ID: 184
    │   │           │   │   └── Node ID: 185
    │   │           │   └── Node ID: 93
    │   │           │       ├── Node ID: 186
    │   │           │       └── Node ID: 187
    │   │           └── Node ID: 47
    │   │               ├── Node ID: 94
    │   │               │   ├── Node ID: 188
    │   │               │   └── Node ID: 189
    │   │               └── Node ID: 95
    │   │                   ├── Node ID: 190
    │   │                   └── Node ID: 191
    │   └── Node ID: 3
    │       └── Node ID: 5
    │           ├── Node ID: 8
    │           │   ├── Node ID: 14
    │           │   │   ├── Node ID: 28
    │           │   │   │   ├── Node ID: 56
    │           │   │   │   │   ├── Node ID: 112
    │           │   │   │   │   │   ├── Node ID: 224
    │           │   │   │   │   │   └── Node ID: 225
    │           │   │   │   │   └── Node ID: 113
    │           │   │   │   │       ├── Node ID: 226
    │           │   │   │   │       └── Node ID: 227
    │           │   │   │   └── Node ID: 57
    │           │   │   │       ├── Node ID: 114
    │           │   │   │       │   ├── Node ID: 228
    │           │   │   │       │   └── Node ID: 229
    │           │   │   │       └── Node ID: 115
    │           │   │   │           ├── Node ID: 230
    │           │   │   │           └── Node ID: 231
    │           │   │   └── Node ID: 29
    │           │   │       ├── Node ID: 58
    │           │   │       │   ├── Node ID: 116
    │           │   │       │   │   ├── Node ID: 232
    │           │   │       │   │   └── Node ID: 233
    │           │   │       │   └── Node ID: 117
    │           │   │       │       ├── Node ID: 234
    │           │   │       │       └── Node ID: 235
    │           │   │       └── Node ID: 59
    │           │   │           ├── Node ID: 118
    │           │   │           │   ├── Node ID: 236
    │           │   │           │   └── Node ID: 237
    │           │   │           └── Node ID: 119
    │           │   │               ├── Node ID: 238
    │           │   │               └── Node ID: 239
    │           │   └── Node ID: 15
    │           │       ├── Node ID: 30
    │           │       │   ├── Node ID: 60
    │           │       │   │   ├── Node ID: 120
    │           │       │   │   │   ├── Node ID: 240
    │           │       │   │   │   └── Node ID: 241
    │           │       │   │   └── Node ID: 121
    │           │       │   │       ├── Node ID: 242
    │           │       │   │       └── Node ID: 243
    │           │       │   └── Node ID: 61
    │           │       │       ├── Node ID: 122
    │           │       │       │   ├── Node ID: 244
    │           │       │       │   └── Node ID: 245
    │           │       │       └── Node ID: 123
    │           │       │           ├── Node ID: 246
    │           │       │           └── Node ID: 247
    │           │       └── Node ID: 31
    │           │           ├── Node ID: 62
    │           │           │   ├── Node ID: 124
    │           │           │   │   ├── Node ID: 248
    │           │           │   │   └── Node ID: 249
    │           │           │   └── Node ID: 125
    │           │           │       ├── Node ID: 250
    │           │           │       └── Node ID: 251
    │           │           └── Node ID: 63
    │           │               ├── Node ID: 126
    │           │               │   ├── Node ID: 252
    │           │               │   └── Node ID: 253
    │           │               └── Node ID: 127
    │           │                   ├── Node ID: 254
    │           │                   └── Node ID: 255
    │           └── Node ID: 9
    │               ├── Node ID: 16
    │               │   ├── Node ID: 32
    │               │   │   ├── Node ID: 64
    │               │   │   │   ├── Node ID: 128
    │               │   │   │   └── Node ID: 129
    │               │   │   └── Node ID: 65
    │               │   │       ├── Node ID: 130
    │               │   │       └── Node ID: 131
    │               │   └── Node ID: 33
    │               │       ├── Node ID: 66
    │               │       │   ├── Node ID: 132
    │               │       │   └── Node ID: 133
    │               │       └── Node ID: 67
    │               │           ├── Node ID: 134
    │               │           └── Node ID: 135
    │               └── Node ID: 17
    │                   ├── Node ID: 34
    │                   │   ├── Node ID: 68
    │                   │   │   ├── Node ID: 136
    │                   │   │   └── Node ID: 137
    │                   │   └── Node ID: 69
    │                   │       ├── Node ID: 138
    │                   │       └── Node ID: 139
    │                   └── Node ID: 35
    │                       ├── Node ID: 70
    │                       │   ├── Node ID: 140
    │                       │   └── Node ID: 141
    │                       └── Node ID: 71
    │                           ├── Node ID: 142
    │                           └── Node ID: 143
    └── Node ID: 2
        ├── Node ID: 7
        │   ├── Node ID: 12
        │   │   ├── Node ID: 24
        │   │   │   ├── Node ID: 48
        │   │   │   │   ├── Node ID: 96
        │   │   │   │   │   ├── Node ID: 192
        │   │   │   │   │   └── Node ID: 193
        │   │   │   │   └── Node ID: 97
        │   │   │   │       ├── Node ID: 194
        │   │   │   │       └── Node ID: 195
        │   │   │   └── Node ID: 49
        │   │   │       ├── Node ID: 98
        │   │   │       │   ├── Node ID: 196
        │   │   │       │   └── Node ID: 197
        │   │   │       └── Node ID: 99
        │   │   │           ├── Node ID: 198
        │   │   │           └── Node ID: 199
        │   │   └── Node ID: 25
        │   │       ├── Node ID: 50
        │   │       │   ├── Node ID: 100
        │   │       │   │   ├── Node ID: 200
        │   │       │   │   └── Node ID: 201
        │   │       │   └── Node ID: 101
        │   │       │       ├── Node ID: 202
        │   │       │       └── Node ID: 203
        │   │       └── Node ID: 51
        │   │           ├── Node ID: 102
        │   │           │   ├── Node ID: 204
        │   │           │   └── Node ID: 205
        │   │           └── Node ID: 103
        │   │               ├── Node ID: 206
        │   │               └── Node ID: 207
        │   └── Node ID: 13
        │       ├── Node ID: 26
        │       │   ├── Node ID: 52
        │       │   │   ├── Node ID: 104
        │       │   │   │   ├── Node ID: 208
        │       │   │   │   └── Node ID: 209
        │       │   │   └── Node ID: 105
        │       │   │       ├── Node ID: 210
        │       │   │       └── Node ID: 211
        │       │   └── Node ID: 53
        │       │       ├── Node ID: 106
        │       │       │   ├── Node ID: 212
        │       │       │   └── Node ID: 213
        │       │       └── Node ID: 107
        │       │           ├── Node ID: 214
        │       │           └── Node ID: 215
        │       └── Node ID: 27
        │           ├── Node ID: 54
        │           │   ├── Node ID: 108
        │           │   │   ├── Node ID: 216
        │           │   │   └── Node ID: 217
        │           │   └── Node ID: 109
        │           │       ├── Node ID: 218
        │           │       └── Node ID: 219
        │           └── Node ID: 55
        │               ├── Node ID: 110
        │               │   ├── Node ID: 220
        │               │   └── Node ID: 221
        │               └── Node ID: 111
        │                   ├── Node ID: 222
        │                   └── Node ID: 223
        └── Node ID: 4
            ├── Node ID: 18
            │   ├── Node ID: 36
            │   │   ├── Node ID: 72
            │   │   │   ├── Node ID: 144
            │   │   │   └── Node ID: 145
            │   │   └── Node ID: 73
            │   │       ├── Node ID: 146
            │   │       └── Node ID: 147
            │   └── Node ID: 37
            │       ├── Node ID: 74
            │       │   ├── Node ID: 148
            │       │   └── Node ID: 149
            │       └── Node ID: 75
            │           ├── Node ID: 150
            │           └── Node ID: 151
            └── Node ID: 19
                ├── Node ID: 38
                │   ├── Node ID: 76
                │   │   ├── Node ID: 152
                │   │   └── Node ID: 153
                │   └── Node ID: 77
                │       ├── Node ID: 154
                │       └── Node ID: 155
                └── Node ID: 39
                    ├── Node ID: 78
                    │   ├── Node ID: 156
                    │   └── Node ID: 157
                    └── Node ID: 79
                        ├── Node ID: 158
                        └── Node ID: 159

#### 320 nodes

└── Node ID: 0
    ├── Node ID: 1
    │   ├── Node ID: 6
    │   │   ├── Node ID: 10
    │   │   │   ├── Node ID: 20
    │   │   │   │   ├── Node ID: 40
    │   │   │   │   │   ├── Node ID: 80
    │   │   │   │   │   │   ├── Node ID: 160
    │   │   │   │   │   │   └── Node ID: 161
    │   │   │   │   │   └── Node ID: 81
    │   │   │   │   │       ├── Node ID: 162
    │   │   │   │   │       └── Node ID: 163
    │   │   │   │   └── Node ID: 41
    │   │   │   │       ├── Node ID: 82
    │   │   │   │       │   ├── Node ID: 164
    │   │   │   │       │   └── Node ID: 165
    │   │   │   │       └── Node ID: 83
    │   │   │   │           ├── Node ID: 166
    │   │   │   │           └── Node ID: 167
    │   │   │   └── Node ID: 21
    │   │   │       ├── Node ID: 42
    │   │   │       │   ├── Node ID: 84
    │   │   │       │   │   ├── Node ID: 168
    │   │   │       │   │   └── Node ID: 169
    │   │   │       │   └── Node ID: 85
    │   │   │       │       ├── Node ID: 170
    │   │   │       │       └── Node ID: 171
    │   │   │       └── Node ID: 43
    │   │   │           ├── Node ID: 86
    │   │   │           │   ├── Node ID: 172
    │   │   │           │   └── Node ID: 173
    │   │   │           └── Node ID: 87
    │   │   │               ├── Node ID: 174
    │   │   │               └── Node ID: 175
    │   │   └── Node ID: 11
    │   │       ├── Node ID: 22
    │   │       │   ├── Node ID: 44
    │   │       │   │   ├── Node ID: 88
    │   │       │   │   │   ├── Node ID: 176
    │   │       │   │   │   └── Node ID: 177
    │   │       │   │   └── Node ID: 89
    │   │       │   │       ├── Node ID: 178
    │   │       │   │       └── Node ID: 179
    │   │       │   └── Node ID: 45
    │   │       │       ├── Node ID: 90
    │   │       │       │   ├── Node ID: 180
    │   │       │       │   └── Node ID: 181
    │   │       │       └── Node ID: 91
    │   │       │           ├── Node ID: 182
    │   │       │           └── Node ID: 183
    │   │       └── Node ID: 23
    │   │           ├── Node ID: 46
    │   │           │   ├── Node ID: 92
    │   │           │   │   ├── Node ID: 184
    │   │           │   │   └── Node ID: 185
    │   │           │   └── Node ID: 93
    │   │           │       ├── Node ID: 186
    │   │           │       └── Node ID: 187
    │   │           └── Node ID: 47
    │   │               ├── Node ID: 94
    │   │               │   ├── Node ID: 188
    │   │               │   └── Node ID: 189
    │   │               └── Node ID: 95
    │   │                   ├── Node ID: 190
    │   │                   └── Node ID: 191
    │   └── Node ID: 3
    │       └── Node ID: 5
    │           ├── Node ID: 8
    │           │   ├── Node ID: 14
    │           │   │   ├── Node ID: 28
    │           │   │   │   ├── Node ID: 56
    │           │   │   │   │   ├── Node ID: 112
    │           │   │   │   │   │   ├── Node ID: 224
    │           │   │   │   │   │   └── Node ID: 225
    │           │   │   │   │   └── Node ID: 113
    │           │   │   │   │       ├── Node ID: 226
    │           │   │   │   │       └── Node ID: 227
    │           │   │   │   └── Node ID: 57
    │           │   │   │       ├── Node ID: 114
    │           │   │   │       │   ├── Node ID: 228
    │           │   │   │       │   └── Node ID: 229
    │           │   │   │       └── Node ID: 115
    │           │   │   │           ├── Node ID: 230
    │           │   │   │           └── Node ID: 231
    │           │   │   └── Node ID: 29
    │           │   │       ├── Node ID: 58
    │           │   │       │   ├── Node ID: 116
    │           │   │       │   │   ├── Node ID: 232
    │           │   │       │   │   └── Node ID: 233
    │           │   │       │   └── Node ID: 117
    │           │   │       │       ├── Node ID: 234
    │           │   │       │       └── Node ID: 235
    │           │   │       └── Node ID: 59
    │           │   │           ├── Node ID: 118
    │           │   │           │   ├── Node ID: 236
    │           │   │           │   └── Node ID: 237
    │           │   │           └── Node ID: 119
    │           │   │               ├── Node ID: 238
    │           │   │               └── Node ID: 239
    │           │   └── Node ID: 15
    │           │       ├── Node ID: 30
    │           │       │   ├── Node ID: 60
    │           │       │   │   ├── Node ID: 120
    │           │       │   │   │   ├── Node ID: 240
    │           │       │   │   │   └── Node ID: 241
    │           │       │   │   └── Node ID: 121
    │           │       │   │       ├── Node ID: 242
    │           │       │   │       └── Node ID: 243
    │           │       │   └── Node ID: 61
    │           │       │       ├── Node ID: 122
    │           │       │       │   ├── Node ID: 244
    │           │       │       │   └── Node ID: 245
    │           │       │       └── Node ID: 123
    │           │       │           ├── Node ID: 246
    │           │       │           └── Node ID: 247
    │           │       └── Node ID: 31
    │           │           ├── Node ID: 62
    │           │           │   ├── Node ID: 124
    │           │           │   │   ├── Node ID: 248
    │           │           │   │   └── Node ID: 249
    │           │           │   └── Node ID: 125
    │           │           │       ├── Node ID: 250
    │           │           │       └── Node ID: 251
    │           │           └── Node ID: 63
    │           │               ├── Node ID: 126
    │           │               │   ├── Node ID: 252
    │           │               │   └── Node ID: 253
    │           │               └── Node ID: 127
    │           │                   ├── Node ID: 254
    │           │                   └── Node ID: 255
    │           └── Node ID: 9
    │               ├── Node ID: 16
    │               │   ├── Node ID: 32
    │               │   │   ├── Node ID: 64
    │               │   │   │   ├── Node ID: 128
    │               │   │   │   │   ├── Node ID: 256
    │               │   │   │   │   └── Node ID: 257
    │               │   │   │   └── Node ID: 129
    │               │   │   │       ├── Node ID: 258
    │               │   │   │       └── Node ID: 259
    │               │   │   └── Node ID: 65
    │               │   │       ├── Node ID: 130
    │               │   │       │   ├── Node ID: 260
    │               │   │       │   └── Node ID: 261
    │               │   │       └── Node ID: 131
    │               │   │           ├── Node ID: 262
    │               │   │           └── Node ID: 263
    │               │   └── Node ID: 33
    │               │       ├── Node ID: 66
    │               │       │   ├── Node ID: 132
    │               │       │   │   ├── Node ID: 264
    │               │       │   │   └── Node ID: 265
    │               │       │   └── Node ID: 133
    │               │       │       ├── Node ID: 266
    │               │       │       └── Node ID: 267
    │               │       └── Node ID: 67
    │               │           ├── Node ID: 134
    │               │           │   ├── Node ID: 268
    │               │           │   └── Node ID: 269
    │               │           └── Node ID: 135
    │               │               ├── Node ID: 270
    │               │               └── Node ID: 271
    │               └── Node ID: 17
    │                   ├── Node ID: 34
    │                   │   ├── Node ID: 68
    │                   │   │   ├── Node ID: 136
    │                   │   │   │   ├── Node ID: 272
    │                   │   │   │   └── Node ID: 273
    │                   │   │   └── Node ID: 137
    │                   │   │       ├── Node ID: 274
    │                   │   │       └── Node ID: 275
    │                   │   └── Node ID: 69
    │                   │       ├── Node ID: 138
    │                   │       │   ├── Node ID: 276
    │                   │       │   └── Node ID: 277
    │                   │       └── Node ID: 139
    │                   │           ├── Node ID: 278
    │                   │           └── Node ID: 279
    │                   └── Node ID: 35
    │                       ├── Node ID: 70
    │                       │   ├── Node ID: 140
    │                       │   │   ├── Node ID: 280
    │                       │   │   └── Node ID: 281
    │                       │   └── Node ID: 141
    │                       │       ├── Node ID: 282
    │                       │       └── Node ID: 283
    │                       └── Node ID: 71
    │                           ├── Node ID: 142
    │                           │   ├── Node ID: 284
    │                           │   └── Node ID: 285
    │                           └── Node ID: 143
    │                               ├── Node ID: 286
    │                               └── Node ID: 287
    └── Node ID: 2
        ├── Node ID: 7
        │   ├── Node ID: 12
        │   │   ├── Node ID: 24
        │   │   │   ├── Node ID: 48
        │   │   │   │   ├── Node ID: 96
        │   │   │   │   │   ├── Node ID: 192
        │   │   │   │   │   └── Node ID: 193
        │   │   │   │   └── Node ID: 97
        │   │   │   │       ├── Node ID: 194
        │   │   │   │       └── Node ID: 195
        │   │   │   └── Node ID: 49
        │   │   │       ├── Node ID: 98
        │   │   │       │   ├── Node ID: 196
        │   │   │       │   └── Node ID: 197
        │   │   │       └── Node ID: 99
        │   │   │           ├── Node ID: 198
        │   │   │           └── Node ID: 199
        │   │   └── Node ID: 25
        │   │       ├── Node ID: 50
        │   │       │   ├── Node ID: 100
        │   │       │   │   ├── Node ID: 200
        │   │       │   │   └── Node ID: 201
        │   │       │   └── Node ID: 101
        │   │       │       ├── Node ID: 202
        │   │       │       └── Node ID: 203
        │   │       └── Node ID: 51
        │   │           ├── Node ID: 102
        │   │           │   ├── Node ID: 204
        │   │           │   └── Node ID: 205
        │   │           └── Node ID: 103
        │   │               ├── Node ID: 206
        │   │               └── Node ID: 207
        │   └── Node ID: 13
        │       ├── Node ID: 26
        │       │   ├── Node ID: 52
        │       │   │   ├── Node ID: 104
        │       │   │   │   ├── Node ID: 208
        │       │   │   │   └── Node ID: 209
        │       │   │   └── Node ID: 105
        │       │   │       ├── Node ID: 210
        │       │   │       └── Node ID: 211
        │       │   └── Node ID: 53
        │       │       ├── Node ID: 106
        │       │       │   ├── Node ID: 212
        │       │       │   └── Node ID: 213
        │       │       └── Node ID: 107
        │       │           ├── Node ID: 214
        │       │           └── Node ID: 215
        │       └── Node ID: 27
        │           ├── Node ID: 54
        │           │   ├── Node ID: 108
        │           │   │   ├── Node ID: 216
        │           │   │   └── Node ID: 217
        │           │   └── Node ID: 109
        │           │       ├── Node ID: 218
        │           │       └── Node ID: 219
        │           └── Node ID: 55
        │               ├── Node ID: 110
        │               │   ├── Node ID: 220
        │               │   └── Node ID: 221
        │               └── Node ID: 111
        │                   ├── Node ID: 222
        │                   └── Node ID: 223
        └── Node ID: 4
            ├── Node ID: 18
            │   ├── Node ID: 36
            │   │   ├── Node ID: 72
            │   │   │   ├── Node ID: 144
            │   │   │   │   ├── Node ID: 288
            │   │   │   │   └── Node ID: 289
            │   │   │   └── Node ID: 145
            │   │   │       ├── Node ID: 290
            │   │   │       └── Node ID: 291
            │   │   └── Node ID: 73
            │   │       ├── Node ID: 146
            │   │       │   ├── Node ID: 292
            │   │       │   └── Node ID: 293
            │   │       └── Node ID: 147
            │   │           ├── Node ID: 294
            │   │           └── Node ID: 295
            │   └── Node ID: 37
            │       ├── Node ID: 74
            │       │   ├── Node ID: 148
            │       │   │   ├── Node ID: 296
            │       │   │   └── Node ID: 297
            │       │   └── Node ID: 149
            │       │       ├── Node ID: 298
            │       │       └── Node ID: 299
            │       └── Node ID: 75
            │           ├── Node ID: 150
            │           │   ├── Node ID: 300
            │           │   └── Node ID: 301
            │           └── Node ID: 151
            │               ├── Node ID: 302
            │               └── Node ID: 303
            └── Node ID: 19
                ├── Node ID: 38
                │   ├── Node ID: 76
                │   │   ├── Node ID: 152
                │   │   │   ├── Node ID: 304
                │   │   │   └── Node ID: 305
                │   │   └── Node ID: 153
                │   │       ├── Node ID: 306
                │   │       └── Node ID: 307
                │   └── Node ID: 77
                │       ├── Node ID: 154
                │       │   ├── Node ID: 308
                │       │   └── Node ID: 309
                │       └── Node ID: 155
                │           ├── Node ID: 310
                │           └── Node ID: 311
                └── Node ID: 39
                    ├── Node ID: 78
                    │   ├── Node ID: 156
                    │   │   ├── Node ID: 312
                    │   │   └── Node ID: 313
                    │   └── Node ID: 157
                    │       ├── Node ID: 314
                    │       └── Node ID: 315
                    └── Node ID: 79
                        ├── Node ID: 158
                        │   ├── Node ID: 316
                        │   └── Node ID: 317
                        └── Node ID: 159
                            ├── Node ID: 318
                            └── Node ID: 319