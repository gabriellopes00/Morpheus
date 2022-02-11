import { Grid, Slider, Typography } from '@mui/material'
import { useFilter } from '../../contexts/event-filter-context'

export function MaximumCapacitySlider() {
  const { filter, updateFilter } = useFilter()

  return (
    <Grid container item>
      <Typography fontWeight="bold">
        Maximum capacity of people: {filter.maximumCapacity}
      </Typography>
      <Slider
        aria-label="maximum-capacity"
        defaultValue={200}
        valueLabelDisplay="auto"
        min={10}
        max={300}
        onChangeCommitted={(_, v) =>
          updateFilter({ ...filter, maximumCapacity: v as number })
        }
      />
    </Grid>
  )
}
