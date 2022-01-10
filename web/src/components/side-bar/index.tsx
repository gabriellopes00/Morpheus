import { Grid } from '@mui/material'
import { AgeGroupSelector } from './age-group-selector'
import { FilterCalendar } from './calendar'
import { MaximumCapacitySlider } from './maximum-capacity-slider'
import { StatesList } from './states-list'

export function SideBar() {
  return (
    <Grid
      container
      gap={5}
      bgcolor={'lightgray'}
      sx={{ overflowY: 'scroll', height: 'calc(100vh - 69px)' }}
      flexDirection="row"
      padding="1rem"
    >
      <FilterCalendar />
      <StatesList />
      <AgeGroupSelector />
      <MaximumCapacitySlider />
    </Grid>
  )
}
