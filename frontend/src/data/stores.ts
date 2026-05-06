export interface Store {
  id: string
  displayName: string
  systemName: string
}

export const STORES: Store[] = [
  { id: 's1', displayName: 'Downtown Plaza', systemName: 'downtown-plaza' },
  { id: 's2', displayName: 'North Bridge Mall', systemName: 'north-bridge-mall' },
  { id: 's3', displayName: 'Riverside Drive', systemName: 'riverside-drive' },
  { id: 's4', displayName: 'Westgate Terminal', systemName: 'westgate-terminal' },
  { id: 's5', displayName: 'Eastfield Park', systemName: 'eastfield-park' },
  { id: 's6', displayName: 'Harborview', systemName: 'harborview' },
  { id: 's7', displayName: 'Maple & 5th', systemName: 'maple-5th' },
  { id: 's8', displayName: 'Airport Concourse C', systemName: 'airport-concourse-c' },
  { id: 's9', displayName: 'University District', systemName: 'university-district' },
  { id: 's10', displayName: 'Southgate Crossing', systemName: 'southgate-crossing' },
]
