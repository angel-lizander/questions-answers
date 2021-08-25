import React, { useState } from 'react';
import { Avatar, CssBaseline, List, ListItem, ListItemAvatar, ListItemText, makeStyles, Paper, Theme, Typography, ListItemSecondaryAction, IconButton, Button, Grid } from "@material-ui/core";
import { useEffect } from 'react';
import { GET, DELETE } from './../services/httpclient/httpclient.services'
import Details from './questionsdetails/Details'
import DeleteIcon from '@material-ui/icons/Delete';
import { Create } from '@material-ui/icons';
import CreateQuestion from './questionscreate/Create'


const useStyles = makeStyles((theme: Theme) =>
({
  text: {
    padding: theme.spacing(2, 2, 0),
  },
  paper: {
    paddingBottom: 50,
  },
  list: {
    marginBottom: theme.spacing(2)
  },
  subheader: {
    backgroundColor: theme.palette.background.paper,
  },
  appBar: {
    top: 'auto',
    bottom: 0,
  },
  grow: {
    flexGrow: 1,
  },
  fabButton: {
    position: 'absolute',
    zIndex: 1,
    top: -30,
    left: 0,
    right: 0,
    margin: '0 auto',
  },
  titlePadding:{
    paddingTop: "12px"
  }
}),
);

function Questions() {

  interface Question {
    id: string,
    QuestionUser: string,
    QuestionTitle: string,
    QuestionDescription: string
  }

  const classes = useStyles();

  const [data, setData] = useState<any | null>(null);
  const [openmodal, setOpenmodal] = useState(false);
  const [id, setId] = useState("");
  const [open, setOpen] = useState(false);


  const GetQuestions = () => {

    GET({ url: 'http://localhost:8080/questions/' }).then(data => setData(data.data))

  }
  const toggleModal = (id: any) => {
    setId(id)
    setOpenmodal(true)
  }

  const onDelete = (id: any) => {

    DELETE({ url: 'http://localhost:8080/questions/' + id }).then(data => data.success ? GetQuestions() : "")

  }

  useEffect(() => {
    GetQuestions()
  }, [data])



  return (
    <div>
      <React.Fragment>
        <CssBaseline />
        <Paper square className={classes.paper}>
          <Grid container className={classes.titlePadding}>
            <Typography className={classes.text} variant="h5" gutterBottom>
              Questions
            </Typography>
            <Button
              color="primary"
              variant="contained"
              endIcon={<Create />}
              onClick={() => setOpen(true)}
            >
              New question
            </Button>
          </Grid>
          <List className={classes.list}>
            {data && data.Questions.map((data: Question) => (
              <React.Fragment key={data.id}>
                <ListItem button onClick={() => toggleModal(data.id)}>
                  <ListItemAvatar>
                    <Avatar alt="Profile Picture" src={'/static/images/avatar/3.jpg'} />
                  </ListItemAvatar>
                  <ListItemText primary={data.QuestionTitle} secondary={data.QuestionDescription} />
                  <ListItemSecondaryAction>
                    <IconButton edge="end" aria-label="delete">
                      <DeleteIcon onClick={() => onDelete(data.id)} />
                    </IconButton>
                  </ListItemSecondaryAction>
                </ListItem>

              </React.Fragment>
            ))}
          </List>
          {openmodal && (
            <Details
              handleClose={() => setOpenmodal(!openmodal)}
              Open={openmodal}
              Id={id}
            />)

          }
          {open &&
            <CreateQuestion
              handleClose={() => setOpen(!open)}
              Open={open}
            />}
        </Paper>
      </React.Fragment>

    </div>
  )
}


export default Questions;