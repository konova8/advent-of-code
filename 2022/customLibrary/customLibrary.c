// LISTE BILINKATE CIRCOLARI

struct Node
{
    struct Node* prev;
    char C;
    struct Node* next;
};

struct Node* inizializeList(struct Node* l, char c)
{
    struct Node* newNode = (struct Node*) malloc(sizeof(struct Node));
    newNode->prev = newNode;
    newNode->C = c;
    newNode->next = newNode;
    return newNode;
}

struct Node* addHead(struct Node* l, char c)
{
    struct Node* newNode = (struct Node*) malloc(sizeof(struct Node));
    newNode->C = c;
    newNode->next = l;
    newNode->prev = l->prev;
    l->prev->next = newNode;
    l->prev = newNode;
    return newNode;
}

struct Node* addTail(struct Node* l, char c)
{
    struct Node* newNode = (struct Node*) malloc(sizeof(struct Node));
    newNode->C = c;
    newNode->next = l;
    newNode->prev = l->prev;
    l->prev->next = newNode;
    l->prev = newNode;
    return l;
}

// Gestire i vari casi
char removeHead(struct Node* l)
{
    char res = l->C;
    l->next->prev = l->prev;
    l->prev->next = l->next;
    free(l);
    return res;
}

// Gestire i vari casi
char removeTail(struct Node* l)
{
    char res = l->prev->C;
    l->prev->prev->next = l;
    l->prev = l->prev->prev;
    return res;
}