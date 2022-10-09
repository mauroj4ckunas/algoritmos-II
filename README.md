Fallos/Errores: 

Primer fallo: Al poner cualquier dni (dni > 0) que no este en el padron, salta error.
Segundo fallo: El defer rompe, printea bien la parte de presidente pero al llegar a gobernador no salta un panic: "panic: runtime error: invalid memory address or nil pointer dereference"