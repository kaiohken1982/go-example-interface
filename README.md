## Interfacce

In questo esercizio creiamo un nuovo tipo, ByteCounter, 
che implementa il metodo Write, utilizzando quindi in maniera intrinseca 
l'interfaccia io.Writer che lo rende compatibile con Fprintf.

Inoltre implementiamo un http middleware per la gestione dei log.


# Utilizzo

Compila 

```
go build .
```

Lancia

```
ADDR=localhost:4033 ./interface
```

# Argomenti trattati

- Interfacce