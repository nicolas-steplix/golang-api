package helpers

/**
 * Coalesce devuelve el valor si no es el valor cero del tipo, de lo contrario devuelve el fallback.
 */
func Coalesce[T comparable](value, fallback T) T {
	var zero T
	if value == zero {
		return fallback
	}
	return value
}
