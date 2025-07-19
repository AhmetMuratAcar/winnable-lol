export function splitByLastHash(str) {
  const index = str.lastIndexOf('#');
  if (index == -1) return [str, ''];
  return [str.slice(0, index), str.slice(index + 1)];
}
