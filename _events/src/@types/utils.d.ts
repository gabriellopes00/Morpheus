/**
 * Modify a type or an interface by adding or changing properties.
 * @example
 * type A = { prop1: string, prop2: string }
 * type B = { prop2: boolean, prop3: string }
 * type C = Modify<A, B> // { prop1: string, prop2: boolean, prop3: string }
 */
export type Modify<T, R> = Pick<T, Exclude<keyof T, keyof R>> & R
