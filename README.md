# Algorythms and data structures


    - ALGORITHMS EXAMINED:

        algorithms/
    ├── comparison_based_sorting/
    │   ├── divide_and_conquer/
    │   │   ├── mergesort/
    │   │   └── quicksort/
    │   ├── simple_comparison_based/
    │   │   ├── insertionsort/
    │   │   ├── selectionsort/
    │   │   ├── bubblesort/
    │   │   └── heapsort/
    ├── linear_time_sorting/
    │   ├── countingsort/
    │   ├── radixsort/
    │   └── bucketsort/
    └── other_algorithms/
        └── factorial/

    - ELEMENTARY DATA STRUCTURES:

        STACK OF ACTIVATION RECORDS
        - stack operations: stack_empty
        - stack operations: push
        - stack operations: pop

        QUEUES
        - priority queues: maximum, extractMax, insert   
        
        - queue operations: accoda
        - queue operations: estrai_da_coda
        - queue operations with checks: init_coda

        LISTS
        - operations on singly linked lists: lista_inserisci_testa
        - operations on singly linked lists: lista_estrai_testa
        - operations on singly linked lists: lista_vuota

        - operations on singly linked lists: lista_inserisci_coda
        - operations on singly linked lists: lista_estrai_coda

        - operations on linked lists: lista_cerca
        - operations on linked lists: lista_cerca_ric
        - operations on linked lists: lista_inserisci_dopo

        - operations on doubly linked lists: listaD_cancella

        - operations on lists with sentinel: lista_cancellaS
        - operations on lists with sentinel: lista_inserisciS_testa
        - operations on lists with sentinel: lista_cercaS

        LISTS WITH INDICES
        - memory allocation: alloca_elemento
        - memory deallocation: dealloca_elemento

    - ORDINAL SELECTION ALGORITHMS:

        - Min Max
        - Rand Select (randpartition)

    - DICTIONARIES AND HASH FUNCTIONS:

        BASIC OPERATIONS
        - search
        - insert
        - delete
        COLLISION RESOLUTION WITH CHAINING
        - search
        - insert
        - delete
        COLLISION RESOLUTION WITH OPEN ADDRESSING
        - search
        - insert
        - delete
        QUADRATIC PROBING
        - verification algorithm

    - TREES:

        DEPTH-FIRST TREE TRAVERSAL
        - Pre Order
        - In Order
        - Post Order
        BREADTH-FIRST TREE TRAVERSAL
        - Breadth

        BINARY SEARCH TREES
        - search: BST_search
        - search: BST_minimum
        - search: BST_maximum
        - search: BST_successor

        - node insertion: BST_insert
        - node deletion: BST_delete

        R&B TREES
        - RB tree rotation: rotate_right
        - RB tree rotation: rotate_left

        - RB tree insertion: RB_insert
        - RB tree deletion: RB_delete
        - RB tree deletion fix: RB_Fix_delete

        B TREES
        -

    - GRAPHS
_________________________________________________________________________________________________________________________
# Folder Descriptions

**comparison_based_sorting**: Contains comparison-based sorting algorithms.
  - **divide_and_conquer**: Algorithms that use the divide and conquer technique.
    - **mergesort**: Implementation and documentation of Merge Sort.
    - **quicksort**: Implementation and documentation of Quick Sort.
  - **simple_comparison_based**: Simple comparison-based sorting algorithms.
    - **insertionsort**: Implementation and documentation of Insertion Sort.
    - **selectionsort**: Implementation and documentation of Selection Sort.
    - **bubblesort**: Implementation and documentation of Bubble Sort.
    - **heapsort**: Implementation and documentation of Heap Sort.

**linear_time_sorting**: Contains sorting algorithms that operate in linear time.
  - **countingsort**: Implementation and documentation of Counting Sort.
  - **radixsort**: Implementation and documentation of Radix Sort.
  - **bucketsort**: Implementation and documentation of Bucket Sort.

**other_algorithms**: Contains other algorithms not classified in the previous categories.
  - **factorial**: Implementation and documentation of the factorial calculation.

**ELEMENTARY DATA STRUCTURES**: Contains implementations of basic data structures.
  - **STACK OF ACTIVATION RECORDS**: Implementations of stack operations.
    - **stack_empty**: Checks if the stack is empty.
    - **push**: Inserts an element into the stack.
    - **pop**: Removes an element from the stack.
  - **QUEUES**: Implementations of queue operations.
    - **maximum**: Finds the maximum element in the priority queue.
    - **extractMax**: Extracts the maximum element from the priority queue.
    - **insert**: Inserts an element into the priority queue.
    - **enqueque**: Inserts an element into the queue.
    - **dequeue**: Removes an element from the queue.
    - **init_queue**: Initializes the queue with checks.

  - **LISTS**: Implementations of list operations.
    - **singly linked lists**: Operations on singly linked lists.
      - **lista_inserisci_testa**: Inserts an element at the head of the list.
      - **lista_estrai_testa**: Extracts an element from the head of the list.
      - **lista_vuota**: Checks if the list is empty.
      - **lista_inserisci_coda**: Inserts an element at the tail of the list.
      - **lista_estrai_coda**: Extracts an element from the tail of the list.
    - **linked lists**: Operations on linked lists.
      - **lista_cerca**: Searches for an element in the list.
      - **lista_cerca_ric**: Recursively searches for an element in the list.
      - **lista_inserisci_dopo**: Inserts an element after a specific node in the list.

**TREE TRAVERSALS**: Contains tree traversal algorithms.
  - **Depth-First Traversal**:
    - **Pre Order**
    - **In Order**
    - **Post Order**
  - **Breadth-First Traversal**:
    - **Level Order**

**BINARY SEARCH TREES**: Contains operations on binary search trees.
  - **search**:
    - **BST_search**: Searches for an element in the BST.
    - **BST_minimum**: Finds the minimum element in the BST.
    - **BST_maximum**: Finds the maximum element in the BST.
    - **BST_successor**: Finds the successor of an element in the BST.
  - **node insertion**:
    - **BST_insert**: Inserts a node into the BST.
  - **node deletion**:
    - **BST_delete**: Deletes a node from the BST.

**R&B TREES**: Contains operations on red-black trees.
  - **RB tree rotation**:
    - **rotate_right**: Performs a right rotation on the RB tree.
    - **rotate_left**: Performs a left rotation on the RB tree.
  - **RB tree insertion**:
    - **RB_insert**: Inserts a node into the RB tree.
  - **RB tree deletion**:
    - **RB_delete**: Deletes a node from the RB tree.
    - **RB_Fix_delete**: Fixes the RB tree after deletion.

**B TREES**: Contains operations on B-trees.
  - **B tree operations**: (To be added)

**GRAPHS**: Contains graph algorithms.

_______________________________________________________________________________________________________________________  
# Algoritmi e strutture dati


    - ALGORITMI ESAMINATI:

    algorithms/
    ├── ordinamento_basato_su_confronti/
    │   ├── divide_et_impera/
    │   │   ├── mergesort/
    │   │   └── quicksort/
    │   ├── ordinamento_semplice_basato_su_confronti/
    │   │   ├── insertionsort/
    │   │   ├── selectionsort/
    │   │   ├── bubblesort/
    │   │   └── heapsort/
    ├── ordinamento_in_tempo_lineare/
    │   ├── countingsort/
    │   ├── radixsort/
    │   └── bucketsort/
    └── altri_algoritmi/
        └── factorial/

    - STRUTTURE DATI ELEMENTARI:

        STACK DEI RECORD DI ATTIVAZIONE
        - operazioni su stack: stack_vuoto
        - operazioni su stack: push
        - operazioni su stack: pop

        CODE
        - code con priorità: maximum, extractMax, insert

        - operazioni su code: accoda
        - operazioni su code: estrai_da_coda
        - operazioni su code con verifiche: init_coda

        LISTE
        - operazioni su liste puntate semplici: lista_estrai_testa
        - operazioni su liste puntate semplici: lista_vuota

        - operazioni su liste puntate semplici: lista_inserisci_coda
        - operazioni su liste puntate semplici: lista_estrai_coda

        - operazioni su liste puntate: lista_cerca
        - operazioni su liste puntate: lista_cerca_ric
        - operazioni su liste puntate: lista_inserisci_dopo

        - operazioni su liste puntate doppie: listaD_cancella

        - operazioni su liste con sentinella: lista_cancellaS
        - operazioni su liste con sentinella: lista_inserisciS_testa
        - operazioni su liste con sentinella: lista_cercaS

        LISTE CON INDICI
        - allocazione memoria: alloca_elemento
        - deallocazione memoria: dealloca_elemento

    - ALGORITMI DI SELEZIONE ORDINALE:

        - Min Max
        - Rand Select (randpartition)

    - DIZIONARI E FUNZIONI HASH:

        OPERAZIONI BASE
        - search
        - insert
        - delete
        RISOLUZIONE COLLISIONI CON CHAINING
        - search
        - insert
        - delete
        RISOLUZIONE COLLISIONI CON INDIRIZZAMENTO APERTO
        - search
        - insert
        - delete
        ISPEZIONE QUADRATICA
        - algoritmo di verifica

    - ALBERI:

        VISITA DI ALBERI IN PROFONDITA'
        - Pre Order
        - In Order
        - Post Order
        VISITA DI ALBERI IN AMPIEZZA
        - Ampiezza

        ALBERI BINARI DI RICERCA
        - ricerca: ABR_ricerca
        - ricerca: ABR_minimo
        - ricerca: ABR_massimo
        - ricerca: ABR_successore
    
        - inserimento di un nodo: ABR_inserisci
        - cancellazione di un nodo: ABR_candella

        ALBERI R&B
        - rotazione alberi RB: rotazione_destra
        - rotazione alberi RB: rotazione_sinistra

        - inserimento in alberi RB: RB_inserisci
        - cancellazione in alberi RB: RB_cancella
        - cancellazione in alberi RB: RB_Fix_cancella

        ALBERI B
        - operazioni elementari: BTree
        - operazioni elementari: Search
        - operazioni elementari: Insert
        - operazioni elementari: Delete

        - procedure ausiliarie: SearchSubtree
        - procedure ausiliarie: SplitChild
        - procedure ausiliarie: InsertNonfull

    - GRAFI:
_________________________________________________________________________________________________________________________
# Descrizione delle Cartelle

**ordinamento_basato_su_confronti**: Contiene algoritmi di ordinamento basati su confronti.
  - **divide_et_impera**: Algoritmi che utilizzano la tecnica divide et impera.
    - **mergesort**: Implementazione e documentazione di Merge Sort.
    - **quicksort**: Implementazione e documentazione di Quick Sort.
  - **ordinamento_semplice_basato_su_confronti**: Algoritmi di ordinamento semplici basati su confronti.
    - **insertionsort**: Implementazione e documentazione di Insertion Sort.
    - **selectionsort**: Implementazione e documentazione di Selection Sort.
    - **bubblesort**: Implementazione e documentazione di Bubble Sort.
    - **heapsort**: Implementazione e documentazione di Heap Sort.

**ordinamento_in_tempo_lineare**: Contiene algoritmi di ordinamento che funzionano in tempo lineare.
  - **countingsort**: Implementazione e documentazione di Counting Sort.
  - **radixsort**: Implementazione e documentazione di Radix Sort.
  - **bucketsort**: Implementazione e documentazione di Bucket Sort.

**altri_algoritmi**: Contiene altri algoritmi non classificati nelle categorie precedenti.
  - **factorial**: Implementazione e documentazione del calcolo del fattoriale.

**STRUTTURE DATI ELEMENTARI**: Contiene implementazioni di strutture dati di base.
  - **STACK DEI RECORD DI ATTIVAZIONE**: Implementazioni delle operazioni su stack.
    - **stack_vuoto**: Verifica se lo stack è vuoto.
    - **push**: Inserisce un elemento nello stack.
    - **pop**: Rimuove un elemento dallo stack.
  - **CODE**: Implementazioni delle operazioni su code.
    - **maximum**: Trova l'elemento massimo nella coda con priorità.
    - **extractMax**: Estrae l'elemento massimo dalla coda con priorità.
    - **insert**: Inserisce un elemento nella coda con priorità.
    - **enqueue**: Inserisce un elemento nella coda.
    - **dequeue**: Rimuove un elemento dalla coda.
    - **init_queue**: Inizializza la coda con controlli.
  - **LISTE**: Implementazioni delle operazioni su liste.
    - **liste semplicemente concatenate**: Operazioni su liste semplicemente concatenate.
      - **lista_inserisci_testa**: Inserisce un elemento in testa alla lista.
      - **lista_estrai_testa**: Estrae un elemento dalla testa della lista.
      - **lista_vuota**: Verifica se la lista è vuota.
      - **lista_inserisci_coda**: Inserisce un elemento in coda alla lista.
      - **lista_estrai_coda**: Estrae un elemento dalla coda della lista.
    - **liste concatenate**: Operazioni su liste concatenate.
      - **lista_cerca**: Cerca un elemento nella lista.
      - **lista_cerca_ric**: Cerca un elemento nella lista ricorsivamente.
      - **lista_inserisci_dopo**: Inserisce un elemento dopo un nodo specifico nella lista.

**ALBERI BINARI DI RICERCA**: Contiene operazioni sugli alberi binari di ricerca.
  - **search**:
    - **BST_search**: Cerca un elemento nell'albero binario di ricerca.
    - **BST_minimum**: Trova l'elemento minimo nell'albero binario di ricerca.
    - **BST_maximum**: Trova l'elemento massimo nell'albero binario di ricerca.
    - **BST_successor**: Trova il successore di un elemento nell'albero binario di ricerca.
  - **node insertion**:
    - **BST_insert**: Inserisce un nodo nell'albero binario di ricerca.
  - **node deletion**:
    - **BST_delete**: Cancella un nodo dall'albero binario di ricerca.

**ALBERI R&B**: Contiene operazioni sugli alberi rosso-neri.
  - **RB tree rotation**:
    - **rotate_right**: Esegue una rotazione a destra sull'albero rosso-nero.
    - **rotate_left**: Esegue una rotazione a sinistra sull'albero rosso-nero.
  - **RB tree insertion**:
    - **RB_insert**: Inserisce un nodo nell'albero rosso-nero.
  - **RB tree deletion**:
    - **RB_delete**: Cancella un nodo dall'albero rosso-nero.
    - **RB_Fix_delete**: Corregge l'albero rosso-nero dopo una cancellazione.

**ALBERI B**: Contiene operazioni sugli alberi B.
  - **B tree operations**: (Da aggiungere)

**GRAFI**: Contiene algoritmi sui grafi.