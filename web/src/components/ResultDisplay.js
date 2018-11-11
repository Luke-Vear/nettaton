import React from 'react'

// no question no correct -- show nothing
// yes quesiton no correct -- show nothing
// yes question yes correct -- show correct
// no question yes correct -- show correct
export const ResultDisplay = ({ correct }) => (
  correct !== null
    ? (
      correct === true ? <div className='result-display correct'>Correct!</div> : <div className='result-display'>Wrong!</div>
    )
    : (
      <div className='result-display' />
    )
)
