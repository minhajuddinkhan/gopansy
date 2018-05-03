CREATE TABLE permitOneForms (
    permitOneId SERIAL,
    isVerified boolean,
    tokenNumber text,
    diaryDate text,
    diaryNumber text,
    importerId text,
    applicantId text,
    nameOfProduct text,
    quantityOfProduct text,
    countOfProduct integer,
    countryOrigin text,
    importedFrom text,
    numberOfConsignment integer,
    exporterName text,
    exporterAddress text,
    destinationPortName text,
    meansOfTransport text,
    portOfEntry text,
    purposeOfImport text,
    departmentUseOnly text,
    PRIMARY KEY (permitOneId)
)

 