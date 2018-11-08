import React from 'react'

// no question no correct -- show nothing
// yes quesiton no correct -- show nothing
// yes question yes correct -- show correct
// no question yes correct -- show correct
export const ResultDisplay = ({ correct }) =>
  correct !== null ? <div>{correctMessage(correct)}</div> : <div />

const correctMessage = (correct) => correct === true ? 'Correct!' : 'Wrong!'
