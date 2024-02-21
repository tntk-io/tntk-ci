import { createTheme } from "@material-ui/core/styles";

export const theme = createTheme({
  palette: {
    primary: {
      light: "#757ce8",
      main: "#117EFF",
      dark: "#1154ff",
      contrastText: "#fff",
    },
    common: {
      black: "#000",
      white: "#fff",
      bg: "#f7f9fc",
      lightBlue: "rgba(222, 230, 240, 0.44)",
      inputbg: "#f6f8fb",
      inputHover: "#edf2fb",
      darkGrey: "#c4c4c4",
      lightGrey: "rgba(30, 32, 45, .05)",
    },
  },
  typography: {
    htmlFontSize: 16,
    fontFamily: "DMSans",
  },
  breakpoints: {
    values: {
      xs: 0,
      sm: 600,
      md: 960,
      lg: 1280,
      xl: 1440,
    },
  },
  overrides: {
    MuiButton: {
      root: {
        height: 46,
        "& $label": {
          fontSize: 14,
          fontWeight: 500,
          textTransform: "none",
        },
      },
      sizeLarge: {
        height: 52,
      },
      outlined: {
        padding: "7px 30px",
        borderRadius: 8,
        "& $label": {
          fontSize: 14,
          fontWeight: 500,
          textTransform: "none",
          color: "#000",
        },
      },
    },
  },
});
