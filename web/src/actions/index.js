export const NEW_QUESTION = 'NEW_QUESTION'
export const ANSWER_QUESTION = 'ANSWER_QUESTION'

export const newQuestion = (question) => ({
  type: 'NEW_QUESTION',
  question
})

export const answerQuestion = () => ({
  type: 'ANSWER_QUESTION'
})

export const fetchQuestion = () => dispatch => {
  return fetch(`https://api.dev.nettaton.com`)
    .then(response => response.json())
    .then(json => dispatch(newQuestion(json)))
}