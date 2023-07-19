Hexagonal (Ports and Adapter Lecture)

1.  Create Repo , Implement Repo / Initial Repo

    1.1 Create Repo Struct

    ```
    type DataRetriverRepo struct {
	    db *gorm.DB
    }
    ```

    1.2 Create Repo Interface
    ```
    type IDataRetriverRepo interface {
	    GetData() (string, error)
    }
    ```

    1.3  Connect Repo with DataSource

    ```
    // Bound GetCompetitor function in to Controller make function able to call GetCompetitor
    func (controller DataRetriverController) GetCompetitors(ctx echo.Context) error {

	    data, err := controller.repository.GetData()

	    if err != nil {
		    return ctx.JSON(http.StatusOK, models.CommonResponse{
			    Code:    2000,
			    Message: "Unable Get Data",
		    })
	    }

	    return ctx.JSON(http.StatusOK, models.CommonResponse{
		    Code:    1000,
		    Message: data,
	    })
    }
    ```

2. Create Controller
    2.1 Create Controller struct which implements the Repository Interface  
    e.g.
    ```
    type DataRetriverController struct {
	    (var) (Irepository)
    }
    ```

    2.2 Create Initializer Function
    // Initializer Function that create NewController that return Address of Controller
    ```
    func NewDataRetriverController(repo repository.DataRetriverRepo) *DataRetriverController {
	return &DataRetriverController{
		repository: repo,
	}
    }
    ```

    2.3 Implement Repository Function
    ```
    // Bound GetCompetitor function in to Controller make function able to call GetCompetitor
    func (controller DataRetriverController) GetCompetitors(ctx echo.Context) error {

	    data, err := controller.repository.GetData()

	    if err != nil {
		    return ctx.JSON(http.StatusOK, models.CommonResponse{
			    Code:    2000,
			    Message: "Unable Get Data",
		    })
	    }

	        return ctx.JSON(http.StatusOK, models.CommonResponse{
		        Code:    1000,
		        Message: data,
	        })
    }
    ```
