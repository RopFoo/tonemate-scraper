export type MusikersuchtFilter = {
  instrument: MusikersuchtInstrument;
  category: MusikersuchtCategory;
};

export enum MusikersuchtInstrument {
  guitar = 4,
  bass = 2,
  drums = 6,
  vocals = 7,
}

export enum MusikersuchtCategory {
  'looking for band' = 'band',
  'looking for musician' = 'musician',
}
