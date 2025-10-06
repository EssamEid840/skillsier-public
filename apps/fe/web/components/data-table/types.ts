export type RowAction<T> = {
  label: string;
  onClick: (row: T) => void | Promise<void>;
};
