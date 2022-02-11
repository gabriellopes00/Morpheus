import { Grid, ToggleButton, ToggleButtonGroup } from '@mui/material'
import { useState } from 'react'
import { useFilter } from '../../contexts/event-filter-context'

export function AgeGroupSelector() {
  const { filter, updateFilter } = useFilter()

  const handleChange = (event: React.MouseEvent<HTMLElement>, age: string) => {
    console.log(age)

    updateFilter({ ...filter, ageGroup: +age })
    console.log(filter)
  }

  return (
    <Grid container item justifyContent="center">
      <ToggleButtonGroup
        value={filter.ageGroup.toString()}
        exclusive
        onChange={handleChange}
      >
        <ToggleButton value="0" color="success">
          L
        </ToggleButton>
        <ToggleButton value="10" color="primary">
          10
        </ToggleButton>
        <ToggleButton value="12" color="warning">
          12
        </ToggleButton>
        <ToggleButton value="14" color="error">
          14
        </ToggleButton>
        <ToggleButton value="16" color="info">
          16
        </ToggleButton>
        <ToggleButton value="18">18</ToggleButton>
      </ToggleButtonGroup>
    </Grid>
  )
}
