const initialState = {
  modal: false,
  component: null,
  modalProps: null
};

export const modalReducer = (state = initialState, action) => {
  switch (action.type) {
    case "OPEN_MODAL":
      return {
        ...state,
        modal: true,
        component: action.payload.component,
        modalProps: action.payload.props || null,
      };

    case "CLOSE_MODAL":
      return { ...state, modal: false, component: null, modalProps: null };

    default:
      return state;
  }
};
