import { createStyles, makeStyles } from "@material-ui/core";

const styles = makeStyles((theme) =>
  createStyles({
    modal: {
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    },
    paper: {
      backgroundColor: theme.palette.background.paper,
      border: '2px solid #000',
      boxShadow: theme.shadows[5],
      padding: theme.spacing(2, 4, 3),
      width: "1000px",
      maxHeight: "85%",
      overflowY: "auto",
      minHeight: "auto"
    },
    container_table: {
      height: "calc(100vh - 260px)",
      marginTop: theme.spacing(3),
      marginBottom: theme.spacing(3),
    },
    formControl: {
      marginBottom: theme.spacing(1),
      marginTop: theme.spacing(1),
      minWidth: 120,
      width: "100%"
    },
    footer: {
      marginTop: theme.spacing(4),
      display: "flex",
      alignItems: "center",
      justifyContent: "center",
    },
    btn_footer: {
      width: "100%"
    },
    textfield: {
      width: "100%"
    },
    title: {
      margin: "20px 8px",
      fontSize: "20px",
      fontWeight: 600
    },
    items: {
      marginTop: "8px",
      fontSize: "13px",
      fontWeight: 600
    },
    loading: {
      display: 'flex',
      '& > * + *': {
        marginLeft: theme.spacing(4),
      }
    },
    heading: {
      fontSize: theme.typography.pxToRem(15),
      fontWeight: theme.typography.fontWeightRegular,
    },

  }),
);

export default styles;