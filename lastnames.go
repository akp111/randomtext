package randomtext

import "math/rand"

func LastName() string {
	var lastNames = []string{"Franklin", "Hernandez", "O'Brien", "Kim", "Gupta", "Morgan", "Zhang", "Mendoza", "Wang", "Riley", "Ali", "Fitzgerald", "Reyes", "Bishop", "Perez", "Sullivan", "Chen", "Graham", "Singh", "Griffin", "O'Connor", "Holland", "Russell", "Gibson", "Torres", "Flynn", "Zimmerman", "Barnes", "Fletcher", "Schneider", "Neal", "Roberts", "Liu", "Rossi", "Murray", "Weaver", "Holmes", "Kumar", "Boyd", "Huang", "Wright", "Moreno", "Ali", "Johnston", "Fleming", "Burns", "Chavez", "Rao","Keller", "Tucker", "Pearson", "McGuire", "Peters", "Ponce", "Schultz", "Sanchez", "Reid", "Baker", "Munoz", "Chan", "Mills", "Duarte", "Watson", "Huerta", "Davies", "Ferguson", "Lambert", "Zhou", "Wang", "Vargas", "Buchanan", "Escobar", "Robinson", "Harris", "Gomez", "Stone", "Holt", "Sheppard", "Fisher", "Blackburn", "Perry", "Clarke", "Bishop", "Richards", "Lowe", "Dickinson", "Bowen", "Hawkins", "Carpenter", "Manning", "Bradshaw", "Carrillo", "Howe", "Farrell", "Walsh", "Henderson", "Gonzalez", "Hale", "Schneider", "Fletcher", "Flores", "Barnett", "Vasquez", "Morrison", "Cunningham", "Gutierrez", "Fleming", "Waters", "Griffin", "Carlson", "Murray", "Nelson", "Gregory", "Foster", "Yu", "Yang", "Patel", "Stevens", "Castillo", "Schmitt", "Jefferson", "Higgins", "Choi", "Erickson", "Lynch", "Whitehead", "Conway", "Lucas", "Nunez", "Wise", "Gould", "Huang", "Stephens", "Wright", "Rivas", "Boyd", "Nicholson", "Garza", "Horne", "Torres", "Zhang", "Gill", "MacDonald", "Lin", "Wade", "Conley", "Wallace", "Roberts", "Hodges", "Garcia", "Mendez", "Wong", "Andrews", "Wu", "Dixon", "Wells", "Ho", "Nash", "Jordan", "Liu", "Hawley", "Kim", "Larson", "Beck", "Hoffman", "Chavez", "Meyer", "Lee", "Hanson", "Wu", "Santos", "Ruiz", "Campbell", "Maldonado", "Rodriguez", "Wilkinson", "Tran", "Sims", "Riley", "Hoover", "Evans", "Morales", "Phillips", "Mendoza", "Hicks", "Taylor", "Harrison", "Owen", "Garcia", "May", "Vaughn", "O'Donnell", "Zimmerman", "Stewart", "Bell", "Burton", "Carson", "Griffiths", "Fox", "Nguyen", "Chung", "Saunders", "Ramirez", "Porter", "Valdez", "Knight", "Sweeney", "O'Neill", "Curtis", "Martinez", "Dawson", "Banks", "Grimes", "Ortega", "Reyes", "Gibson", "Butler", "Petersen", "Hopkins", "Blair", "Duncan", "Silva", "Gutierrez", "Gill", "Abbott", "Cox", "Howell", "Roy", "Boyle", "Carney", "Floyd", "Salazar", "Stanton", "Warner", "Hendrix", "Dominguez", "McCarthy", "Dunn", "Woodward", "Sandoval", "Lopez", "Pittman", "Holmes", "Alexander", "Mckinney", "Olsen", "Ramos", "Yu", "Lara", "Pacheco", "Bowers", "Stokes", "Morris", "Brown", "Ford", "Gaines", "Mckay", "West", "Lloyd", "Schroeder", "Mcclain", "Pugh", "Zhu", "Beasley", "Mclaughlin", "Horton", "Barrett", "Barker", "Byrd", "Yang", "Thornton", "Murphy", "Jacobs", "Collins", "Rice", "Mendez", "Simpson", "Dickson", "Osborne", "Mercado", "Nguyen", "Haynes", "Sharp", "Bryant", "Sanchez", "Davidson", "Peters", "Nichols", "Ferguson", "Brewer", "Hancock", "Richardson", "Gibbs", "Washington", "Hutchinson", "Rivera", "Brennan", "Parrish", "Morton", "Martin", "Guerra", "Benson", "Fields", "Montgomery", "Mcdonald", "Dalton", "Huffman", "Hardy", "Li", "Warren", "Gordon", "Bolton", "Simon", "Snow", "Barber", "Colon", "Morrow", "Blake", "Sullivan", "Bush", "Chandler", "Johnston", "Cohen", "Powell", "Hudson", "Mann", "Harrison", "Casey", "Pierce", "Brady", "Moss", "Marshall", "Saunders", "Paul", "Carr", "Stephenson", "Gomez", "Burgess", "Burns", "Perez", "Carroll", "Knox", "Pitts", "Scott", "Townsend", "Frost", "Mata", "Barry", "Marquez", "Le", "Franco", "Barrera", "Roberts", "Goodman", "Schmidt", "Bates", "Wise", "Fields", "Villarreal", "Woodward", "Grant", "Howard", "Terry", "Bentley", "Garrett", "Holloway", "Greer", "Bowers", "Yu", "Hensley", "Leblanc", "Parrish", "Singleton", "Mcclure", "Pugh", "Mcmillan", "Wolfe", "Olsen", "Brewer", "Mccarthy", "Molina", "Rice", "Bass", "Blanchard", "Foley", "Le", "Hardin", "Gill", "Trujillo", "Ray", "Mcconnell", "Ho", "Faulkner", "Hess", "Harrington", "Herman", "Baldwin", "Collier", "Duran", "Horton", "Owens", "Zhang", "Pittman", "Cantu", "Hess", "Bennett", "Morton", "Hines", "Guerra", "Case", "Cantu", "Reilly", "Dickerson", "Bullock", "Werner", "Huang", "Mueller"}
	randomIndex := rand.Intn(len(lastNames))
	return lastNames[randomIndex]
}