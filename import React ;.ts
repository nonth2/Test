import React ;
interface Trashbin {
Id: string;
Location: string;
fillLevel: number;
}
async function fetchTrashbin(): Promise<Trashbin[]> {
    const response = await fetch('api/trashbin');
    const data = await response.json();
    return data;
}
function TrashBinDashboard(){
    const [Trashbin, setTrashbin] = useState<Trashbin[]>([]);
    useEffect(() => {
      fetchTrashbin().then((data)=>setTrashbin(data));
      }
    }
return (
    <div>
        {users.map((user) => (
            <div key={Trashbin.ID}>
                <h2>{Trashbin.Location}</h2>
                <p>{Trashbin.fillLevel}</p>
        ))}
)
