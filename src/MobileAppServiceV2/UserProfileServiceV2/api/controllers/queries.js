exports.INSERT_USER_PROFILE =
'INSERT INTO userprofiles\
(\
Id,\
FirstName,\
LastName,\
UserId,\
ProfilePictureUri,\
Rating,\
Ranking,\
TotalDistance,\
TotalTrips,\
TotalTime,\
HardStops,\
HardAccelerations,\
FuelConsumption,\
MaxSpeed,\
CreatedAt,\
UpdatedAt,\
Deleted\
) \
SELECT \
Id,\
FirstName,\
LastName,\
UserId,\
ProfilePictureUri,\
Rating,\
Ranking,\
TotalDistance,\
TotalTrips,\
TotalTime,\
HardStops,\
HardAccelerations,\
FuelConsumption,\
MaxSpeed,\
CreatedAt,\
UpdatedAt,\
Deleted \
FROM OPENJSON(@UserProfileJson) \
WITH \
 ( \
    Id nvarchar(128),\
	FirstName nvarchar(max),\
	LastName nvarchar(max),\
	UserId nvarchar(max),\
	ProfilePictureUri nvarchar(max),\
	Rating int,\
	Ranking int,\
	TotalDistance float(53),\
	TotalTrips bigint,\
	TotalTime bigint,\
	HardStops bigint,\
	HardAccelerations bigint,\
	FuelConsumption float(53),\
	MaxSpeed float(53),\
	CreatedAt datetimeoffset,\
	UpdatedAt datetimeoffset,\
	Deleted bit\
 ) as json'

 exports.SELECT_USER_PROFILE_BY_ID=
 'select * from userprofiles where id = @user_profile_id for json path, without_array_wrapper'

 exports.SELECT_USER_PROFILES=
 'select * FROM userprofiles FOR JSON PATH, ROOT(\'user_profiles\')'

 exports.DELETE_USER_PROFILES