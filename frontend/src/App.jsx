import {useEffect, useState} from 'react';
import logo from './assets/images/logo-universal.png';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus, faTrash, faCopy, faPen } from '@fortawesome/free-solid-svg-icons';
import './App.css';
import {Greet, GetDB, AddNoteToDB, DeleteNoteFromDB, EditNoteToDB, GetKey} from "../wailsjs/go/main/App";
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    Modal,
    TextField
} from "@mui/material";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);
    const [modalIsOpen, setModalIsOpen] = useState(false);
    const [modalEditIsOpen, setModalEditIsOpen] = useState(false);
    const [title, setTitle] = useState('Ваше название');
    const [content, setContent] = useState('Логин, пароль и что-то еще...');
    const [titleEdit, setTitleEdit] = useState('Ваше название');
    const [contentEdit, setContentEdit] = useState('Логин, пароль и что-то еще...');
    const [key, setKey] = useState('AES-256-CBC');
    const [modalKeyIsOpen, setModalKeyIsOpen] = useState(true);

    const [notes, setNotes] = useState([]);

    const [selectedNote, setSelectedNote] = useState(null);
    const [dividerPosition, setDividerPosition] = useState('40%');
    const minDividerPosition = '30%';
    const maxDividerPosition = '70%';

    useEffect(() => {
        //getDB();
    }, []);

    window.addEventListener('contextmenu', function (event) {
        event.preventDefault();
    });


    const handleNoteClick = (note) => {
        setSelectedNote(note);
    };

    const handleDividerDrag = (event) => {
        const newPosition = `${(event.clientX / window.innerWidth) * 100}%`;
        // Проверяем, чтобы новая позиция не выходила за границы
        if (newPosition >= minDividerPosition && newPosition <= maxDividerPosition) {
            setDividerPosition(newPosition);
        }
    };

    const handleClose = () => {
        setModalIsOpen(false);
        setModalEditIsOpen(false);
        setModalKeyIsOpen(false)
    };

    // Function to handle form submission
    const handleSubmit = async (event) => {
        event.preventDefault();
        await addNote()
        await handleClose();
        await restartNotes()
    };

    const handleEditSubmit = async (event) => {
        event.preventDefault();
        console.log(titleEdit, contentEdit)
        await editNote()
        await handleClose();
        await restartNotes()
    }

    const handleKeySubmit = async (event) => {
        event.preventDefault();
        await GetKey(key).then(async b => {
            if (b === true) {
                await handleClose();
                await getDB();
            } else {
                await window.close();
            }
        });
    }

    const restartNotes = () => {
        setNotes([])
        getDB()
    }

    const deleteNote = async () => {
        await DeleteNoteFromDB(selectedNote.id).then()
        setSelectedNote(null)
        await restartNotes()
    }

    const copyNote = () => {
        navigator.clipboard.writeText(selectedNote.content).then()
    }

    const openModal = () => {
        setTitle("Ваше название");
        setContent("Логин, пароль и что-то еще...");
        setModalIsOpen(true);
    }

    const openEditModal = () => {
        setTitleEdit(selectedNote.title);
        setContentEdit(selectedNote.content);
        setModalEditIsOpen(true);
    }



    function greet() {
        Greet(name).then(updateResultText);
    }

    function getDB() {
        GetDB().then(data => {
            const encoder = new TextEncoder();
            const uint8Array = encoder.encode(data); // Convert string to Uint8Array
            const decoder = new TextDecoder();
            const jsonData = JSON.parse(decoder.decode(uint8Array));
            console.log(jsonData);

            // Преобразуем массивы в объект
            const formattedData = {
                id: jsonData.id.map((id, index) => ({ id, title: jsonData.title[index], content: jsonData.content[index] }))
            };

            console.log(formattedData);

            setNotes([formattedData]); // Сохраняем отформатированные данные
        })
    }

    function addNote() {
        const jsonNote = {
            title: title,
            content: content,
        };
        console.log(jsonNote)
        let jsonData = JSON.stringify(jsonNote).toString();
        console.log(jsonData);
        AddNoteToDB(jsonData).then();
        setTitle("");
        setContent("");
    }

    function editNote() {
        EditNoteToDB(selectedNote.id, titleEdit, contentEdit).then();
    }

    return (
        <div id="App">
            <div className="app-container">
                <div className="top-menu">
                    <div className="menu-item" onClick={() => openModal()}>
                        <FontAwesomeIcon icon={faPlus}/>
                    </div>
                    <div className="menu-item" onClick={() => deleteNote()}>
                        <FontAwesomeIcon icon={faTrash}/>
                    </div>
                    <div className="menu-item" onClick={() => copyNote()}>
                        <FontAwesomeIcon icon={faCopy}/>
                    </div>
                    <div className="menu-item" onClick={() => openEditModal()}>
                        <FontAwesomeIcon icon={faPen}/>
                    </div>
                </div>
                <div className="sidebar" style={{width: dividerPosition}}>
                    <h2 style={{userSelect: 'none'}}>Заметки</h2>
                    <ul style={{userSelect: 'none'}}>
                        {notes.map((noteGroup, index) => (
                            Object.keys(noteGroup).map((key) => (
                                noteGroup[key].map((note, idx) => (
                                    <li key={note.id} onClick={() => handleNoteClick(note)}>
                                        {note.title}
                                    </li>
                                ))
                            ))
                        ))}
                    </ul>
                </div>

                <div
                    className="divider"
                    style={{left: dividerPosition, transform: 'translateX(-50%)'}}
                    onMouseDown={() => {
                        document.addEventListener('mousemove', handleDividerDrag);
                        document.addEventListener('mouseup', () => {
                            document.removeEventListener('mousemove', handleDividerDrag);
                        });
                    }}
                />
                <div className="main-content" style={{width: `calc(100% - ${dividerPosition})`}}>
                    {selectedNote ? (
                        <div>
                            <h2 style={{ userSelect: 'none' }}>{selectedNote.title}</h2>
                            <p style={{ wordWrap: 'break-word' }}>{selectedNote.content}</p>
                        </div>
                    ) : (
                        <div>
                            <h2>Выберите заметку</h2>
                            <p style={{ color: 'black', fontSize: 14 }}>Надежное хранилище для ваших секретов. Приложение шифрует данные и хранит локально.</p>
                        </div>
                    )}
                </div>
                <Dialog open={modalIsOpen} onClose={handleClose}>
                    <DialogTitle style={{userSelect: 'none'}}>Новая запись</DialogTitle>
                    <DialogContent>
                        <DialogContentText style={{userSelect: 'none'}}>
                            Здесь вы можете создать новый секрет, который
                            никто не узнает. Приложение использует шифрование и хранит данные локально.
                        </DialogContentText>
                        <form onSubmit={handleSubmit}>
                            <TextField
                                autoFocus
                                required
                                margin="dense"
                                id="title"
                                name="title"
                                label="Название записи"
                                fullWidth
                                variant="standard"
                                defaultValue={title}
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                            />
                            <div style={{userSelect: 'none'}}>|</div>
                            <TextField
                                required
                                id="content"
                                name="content"
                                label="Запись"
                                multiline
                                rows={1}
                                fullWidth
                                defaultValue={content}
                                value={content}
                                onChange={(e) => setContent(e.target.value)}
                            />
                            <DialogActions>
                                <Button onClick={handleClose}>Отменить</Button>
                                <Button type="submit">Добавить</Button>
                            </DialogActions>
                        </form>
                    </DialogContent>
                </Dialog>
                <Dialog open={modalEditIsOpen} onClose={handleClose}>
                    <DialogTitle style={{userSelect: 'none'}}>Изменить запись</DialogTitle>
                    <DialogContent>
                        <DialogContentText style={{userSelect: 'none'}}>
                            Измените название или содержание.
                        </DialogContentText>
                        <form onSubmit={handleEditSubmit}>
                            <TextField
                                autoFocus
                                required
                                margin="dense"
                                id="title"
                                name="title"
                                label="Название записи"
                                fullWidth
                                variant="standard"
                                defaultValue={titleEdit}
                                value={titleEdit}
                                onChange={(e) => setTitleEdit(e.target.value)}
                            />
                            <div style={{userSelect: 'none'}}>|</div>
                            <TextField
                                required
                                id="content"
                                name="content"
                                label="Запись"
                                multiline
                                rows={1}
                                fullWidth
                                defaultValue={contentEdit}
                                value={contentEdit}
                                onChange={(e) => setContentEdit(e.target.value)}
                            />
                            <DialogActions>
                                <Button onClick={handleClose}>Отменить</Button>
                                <Button type="submit">Изменить</Button>
                            </DialogActions>
                        </form>
                    </DialogContent>
                </Dialog>
                <Dialog open={modalKeyIsOpen} onClose={handleClose}>
                    <DialogTitle style={{userSelect: 'none'}}>Введите ключ</DialogTitle>
                    <DialogContent>
                        <DialogContentText style={{userSelect: 'none'}}>
                            Для открытия приложения требется ввести ключ.
                        </DialogContentText>
                        <form onSubmit={handleKeySubmit}>
                            <TextField
                                autoFocus
                                required
                                margin="dense"
                                id="key"
                                name="keys"
                                label="Ключ"
                                fullWidth
                                variant="standard"
                                defaultValue={key}
                                value={key}
                                onChange={(e) => setKey(e.target.value)}
                            />
                            <div style={{userSelect: 'none'}}>|</div>
                            <DialogActions>
                                <Button type="submit">Открыть</Button>
                            </DialogActions>
                        </form>
                    </DialogContent>
                </Dialog>
            </div>
        </div>
    )
}

export default App
