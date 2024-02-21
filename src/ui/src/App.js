import React from 'react';
import AppNavigator from './components/app-navigator';
import { ThemeProvider } from "@material-ui/core/styles";
import Modal from './components/UI/modal';
import { theme } from "./helpers/theme";

const App = () => {
  return (
    <ThemeProvider theme={theme}>
      <AppNavigator />
      <Modal />
    </ThemeProvider>
  );
};

export default App;