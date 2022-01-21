import { Autocomplete } from '@mui/material'
import { SearchInput } from './search-input'

export function SearchBox() {
  return (
    <Autocomplete
      handleHomeEndKeys
      clearOnEscape
      id="combo-box-demo"
      options={top100Films.map((option) => option.label)}
      sx={{ width: 300 }}
      renderInput={(params) => <SearchInput params={params} />}
    />
  )
}

const top100Films = [
  { label: 'The Shawshank Redemption', year: 1994 },
  { label: 'The Godfather', year: 1972 },
  { label: 'The Godfather: Part II', year: 1974 },
  { label: 'The Dark Knight', year: 2008 },
  { label: '12 Angry Men', year: 1957 },
  { label: "Schindler's List", year: 1993 },
  { label: 'Pulp Fiction', year: 1994 },
]
