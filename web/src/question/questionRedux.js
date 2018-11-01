export const types = {
  API_CALL_REQUEST: 'QUESTION/API_CALL_REQUEST',
  API_CALL_SUCCESS: 'QUESTION/API_CALL_SUCCESS',
  API_CALL_FAILURE: 'QUESTION/API_CALL_FAILURE'
}

export const actions = {
  apiCallRequest: () => ({ type: types.API_CALL_REQUEST }),
  apiCallSuccess: question => ({ type: types.API_CALL_SUCCESS, question }),
  apiCallFailure: error => ({ type: types.API_CALL_FAILURE, error })
}

const initialState = {
  fetching: false,
  question: null,
  error: null
}

export function reducer (state = initialState, action) {
  switch (action.type) {
    case types.API_CALL_REQUEST:
      return { ...state, fetching: true, error: null }
    case types.API_CALL_SUCCESS:
      return { ...state, fetching: false, question: action.question }
    case types.API_CALL_FAILURE:
      return { ...state, fetching: false, question: null, error: action.error }
    default:
      return state
  }
}
