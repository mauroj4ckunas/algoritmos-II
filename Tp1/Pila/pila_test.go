package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.NotNil(t, pila)
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaOrdenada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(5)
	pila.Apilar(10)
	pila.Apilar(7)
	require.EqualValues(t, 7, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 7, pila.Desapilar())
	require.EqualValues(t, 10, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 10, pila.Desapilar())
	require.EqualValues(t, 5, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 5, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= 100000; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	for i := 100000; i >= 0; i-- {
		require.EqualValues(t, i, pila.VerTope())
		require.False(t, pila.EstaVacia())
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	pila.Apilar(123)
	pila.Apilar(252)
	pila.Apilar(35)
	pila.Apilar(89)
	pila.Apilar(1)
	require.EqualValues(t, 1, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 1, pila.Desapilar())
	require.EqualValues(t, 89, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 89, pila.Desapilar())
	require.EqualValues(t, 35, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 35, pila.Desapilar())
	require.EqualValues(t, 252, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 252, pila.Desapilar())
	require.EqualValues(t, 123, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 123, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}
