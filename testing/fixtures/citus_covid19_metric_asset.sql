insert into covid19.metric_asset (id, label, body, last_modified, released)
values  ('97c6dd89-cf94-4e48-b88b-4847ea839e67', 'Abstract new cases change by date reported', 'The difference between the number of new cases (people who have had at least one positive COVID-19 test result) during the latest 7-day period and the number for the previous, non-overlapping, 7-day period. Data are shown by the date the figures appeared in the published totals.', '2021-09-02 15:14:02.415920', false),
        ('87665f61-e9b8-438b-a861-c518221d0287', 'Abstract new cases LFD only rolling rate by specimen date', 'Rate per 100,000 people of the number of new cases (people who have had at least one positive COVID-19 test result) that were identified by a positive rapid lateral flow (LFD) test and were not confirmed by a positive polymerase chain reaction (PCR) test taken within 3 days, within rolling 7-day periods. Data are shown by the date the lateral flow test was taken.', '2021-09-02 15:27:36.526867', false),
        ('c9a59f4d-a2ca-4d70-8a27-183e1031e145', 'Abstract new deaths within 28 days of a positive test rolling sum by publish date', 'The number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test within rolling 7-day periods. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 13:15:30.516145', false),
        ('45e8eed5-8ebe-44ac-bd2a-27cb605bf474', 'Abstract new deaths within 60 days of a positive test rolling rate by death date', 'Rate per 100,000 people of the number of deaths of people who had a positive test for COVID-19 and either died within 60 days of their first positive test or have COVID-19 mentioned on their death certificate within rolling 7-day periods. Data are shown by the dates the deaths occurred.', '2021-09-03 13:23:04.770854', false),
        ('9b8760f3-4769-447f-becc-84f48117214e', 'Abstract new pillars 1 and 2 tests by publish date', 'Daily numbers of new confirmed positive, negative or void COVID-19 tests conducted under pillar 1 – NHS and PHE COVID-19 testing and pillar 2 – the UK Government COVID-19 testing programme. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:56:39.881599', false),
        ('a5ead37c-9b02-4812-bc81-f61b6bf5defa', 'Cases by Lower Super Output Area (LSOA)', 'Weekly number of cases time series data are available to download for English LSOAs via the following links. The data are updated on Thursdays. Counts between 0 and 2 are denoted by -99.

- [Weekly COVID-19 cases by LSOA in CSV format](https://coronavirus.data.gov.uk/downloads/lsoa_data/LSOAs_latest.csv)
- [Weekly COVID-19 cases by LSOA in JSON format](https://coronavirus.data.gov.uk/downloads/lsoa_data/LSOAs_latest.json)
- [Weekly COVID-19 cases by LSOA in XML format](https://coronavirus.data.gov.uk/downloads/lsoa_data/LSOAs_latest.xml)', '2021-06-23 15:21:26.577223', false),
        ('e07bfa94-f1e2-4e4b-988b-9da7f0a2323c', 'Cumulative total number of deaths within 28 days of positive test, by age and sex', 'Total number of deaths since the start of the pandemic of people who had had a positive test result for COVID-19 and died with 28 days of the first positive test, and death rates per 100,000 resident population. Some records have missing age or sex, so the sum of the subgroups does not equal the total deaths for the area.', '2021-06-23 10:00:06.951027', false),
        ('9aa28e42-a57c-4187-a3ba-c751f0baf746', 'Cumulative number of patients admitted to hospital, by age', 'Total number of patients admitted to hospital with COVID-19 since the start of the pandemic, by age. Some admission records have missing age, so the sum of the subgroups does not equal the total admissions for the area.', '2021-06-24 09:32:48.116087', false),
        ('7be0bc08-05bb-4647-901e-105c622374fc', 'Abstract cumulative case rate by date reported', 'Rate per 100,000 people of the total number of cases (people who have had at least one positive COVID-19 test result) since the start of the pandemic.
Data are shown by the date the figures appeared in the published totals.', '2021-09-02 15:12:53.000826', false),
        ('9e2fa975-6c80-4a42-8a3d-a19dc685bdab', 'Abstract new cases LFD only rolling sum by specimen date', 'The number of new cases (people who have had at least one positive COVID-19 test result) that were identified by a positive rapid lateral flow (LFD) test and were not confirmed by a positive polymerase chain reaction (PCR) test taken within 3 days, within rolling 7-day periods. Data are shown by the date the lateral flow test was taken.', '2021-09-02 15:32:59.926039', false),
        ('50f35533-ebd7-4f82-8901-922e8000cf92', 'Abstract new deaths within 60 days of a positive test rolling sum by death date', 'The number of deaths of people who had a positive test for COVID-19 and either died within 60 days of their first positive test or have COVID-19 mentioned on their death certificate within rolling 7-day periods.  Data are shown by the dates the deaths occurred.', '2021-09-03 13:25:28.259882', false),
        ('fcda3664-5145-4f45-a622-a7bfbeeae70f', 'Abstract new deaths within 60 days of a positive test by publish date', 'Daily numbers of deaths of people who had a positive test for COVID-19 and either died within 60 days of their first positive test or have COVID-19 mentioned on their death certificate. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 13:29:19.538245', false),
        ('943c0631-ded2-49f4-827b-29a16f8b9ca8', 'Abstract new pillar 3 tests by publish date', 'Daily numbers of new confirmed positive, negative or void COVID-19 antibody tests conducted under pillar 3 – antibody serology testing. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:58:15.742241', false),
        ('374011d1-48e8-475e-b7eb-61344f47fe13', 'Deaths long (need to check what to call this)', 'The data do not include deaths of people who had COVID-19 but had not been tested or people who had been tested negative and subsequently caught the virus and died.

Deaths of people who have tested positively for COVID-19 could in some cases be due to a different cause.

# England
Data on COVID-19 deaths in England are produced by Public Health England (PHE). These data are taken from 3 different sources:

NHS England: deaths are reported by NHS Trusts using the [COVID-19 Patient Notification System (CPNS)](https://www.england.nhs.uk/coronavirus/wp-content/uploads/sites/52/2020/04/C0389-update-to-cpns-reporting-letter-27-april-2020.pdf) (this includes only deaths in hospitals)

PHE Health Protection Teams: the local teams report deaths notified to them (mainly deaths not in hospitals)

Linking data on confirmed positive cases (identified through testing by NHS and PHE laboratories and commercial partners) to the NHS Demographic Batch Service: when a patient dies, the NHS central register of patients is notified (this is not limited to deaths in hospitals). The list of all lab-confirmed cases is checked against the NHS central register each day, to check if any of the patients have died.

Data on deaths from these 3 sources are linked to the list of people who have had a diagnosis of COVID-19 confirmed by a PHE or NHS laboratory (pillar 1 of the government''s mass testing strategy) or through testing in commercial labs (pillar 2). This is to identify as many people with a confirmed diagnosis who have died as possible.

Deaths will often appear in 2 or 3 different sources, so the records are checked and merged into one database and duplicates are removed so there is no double counting. Automated processes are used to ensure that the data are as complete as possible. Full details of the process of producing the data are available on [GOV.UK.](https://www.gov.uk/government/publications/phe-data-series-on-deaths-in-people-with-covid-19-technical-summary)

The final list of deaths includes all deaths previously reported by [NHS England](https://www.england.nhs.uk/statistics/statistical-work-areas/covid-19-daily-deaths/), but also includes other deaths of patients who were confirmed cases, whether they died in hospital or elsewhere.

Deaths reported for England each day cover the 24 hours up to 5pm on the day before the reporting date stated above.

# Northern Ireland
Data for Northern Ireland include all cases reported to the Public Health Agency (PHA) where the deceased had a positive test for COVID-19 and died within 28 days, whether or not COVID-19 was the cause of death. PHA sources include reports by healthcare workers (eg Health and Social Care Trusts, GPs) and information from local laboratory reports. Deaths reported each day cover the 24 hours up to 9:30am on the day before the reporting date stated above. The deaths are reported a day earlier by the Northern Ireland Department of Health.

# Scotland
Data for Scotland include deaths which have been registered with National Records of Scotland (NRS) where a laboratory confirmed report of COVID-19 in the 28 days prior to death exists.

The daily total is an update of the data described above/previously, using the latest daily information received from NRS to check where a laboratory positive report for COVID-19 exists. These data include all deaths in individuals with laboratory confirmed COVID-19 in Scotland and therefore include deaths in hospitals, care homes and the community.   This daily number of new deaths registered will not necessarily equal the number of deaths on a particular day, owing to the time allowed for families to register deaths. Numbers of registrations are also lower at weekends.

Deaths reported each day cover the 24 hours up to 9:30am on the day before the reporting date stated above. The deaths are reported a day earlier by the Scottish Government.

# Wales
Data for Wales include reports to Public Health Wales of deaths of hospitalised patients in Welsh Hospitals or care home residents where COVID-19 has been confirmed with a positive laboratory test and the clinician suspects this was a causative factor in the death.

The figures do not include individuals who may have died from COVID-19 but who were not confirmed by laboratory testing, those who died in other settings, or Welsh residents who died outside of Wales.

Deaths reported each day cover the 24 hours up to 5pm on the day before the reporting date stated above.', '2021-06-23 14:30:10.218076', false),
        ('34d42a90-2d12-42a1-ac7d-0a23642375b6', 'Daily deaths with COVID-19 on the death certificate, by date of registration', 'Total number of deaths of people whose death certificate mentioned COVID-19 as one of the causes, registered each week.', '2021-06-24 14:35:00.740384', false),
        ('7feb002e-968a-44db-aafa-34c97dddcb4e', 'Daily mechanical ventilation bed occupancy', 'Mechanical ventilation bed occupancy by occupant type.', '2021-06-24 09:31:23.295737', false),
        ('ea36ecbc-8d6e-4efb-8f4f-0442bd50cea7', 'Case rate by age group', 'Rate of people with a positive COVID-19 virus test result per 100,000 population by specimen date, broken down by 5-year age group.', '2021-06-24 14:55:15.464196', false),
        ('2bacb437-a1b4-4f3b-879f-5cafea49dc9c', 'Daily COVID-19 vaccinations given, by vaccination date', 'Number of people who have received a COVID-19 vaccination, by vaccination date.', '2021-06-24 09:21:15.036605', false),
        ('cf779c35-e657-420a-bebf-4dbef9a5b5a7', 'Abstract cumulative cases by specimen date', 'Total number of cases (people who have had at least one positive COVID-19 test result) since the start of the pandemic. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:17:50.462955', false),
        ('5bfc2aec-43ac-49b2-b389-9acf203e36f5', 'Abstract new cases PCR only by specimen date', 'Daily numbers of new cases (people who have had at least one positive COVID-19 test result) that were identified by a positive polymerase chain reaction (PCR) test, excluding people who had a positive rapid lateral flow (LFD) test within 3 days before the positive PCR test. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:33:57.833987', false),
        ('ce7aed01-2feb-410e-9ad4-8772a8f2b18e', 'Abstract cumulative deaths within 28 days of a positive test by death date', 'Total number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test since the start of the pandemic. Data are shown by the dates the deaths occurred.', '2021-09-03 13:35:32.376257', false),
        ('52799d73-fd01-4fdc-aa9b-3463d53745f2', 'Abstract new pillar 2 tests by publish date', 'Daily numbers of new confirmed positive, negative or void COVID-19 tests conducted under pillar 2 – the UK Government COVID-19 testing programme. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:59:54.471422', false),
        ('93fdc3b7-516c-4755-bd3e-3ccb76f6527b', 'Cases by Middle Super Output Area (MSOA)', 'Weekly rolling sums and population-based rates of new cases by specimen date time series data are available to download for English MSOAs via the dynamic selections on the download data page. The data are updated each day, and show the latest 7 days for which near-complete data — release date minus 5 days — are available, and historic non-overlapping 7-day blocks. Dates are the final day in the relevant 7-day block, and counts between 0 and 2 are blank in the CSV or NULL in the other formats.', '2021-06-30 10:19:26.121065', false),
        ('9ac650bd-d0a5-4c79-b79b-c2a5f0dec2d1', 'Daily COVID-19 patients in mechanical ventilation beds', 'COVID-19 patients in mechanical ventilation beds.', '2021-06-24 09:35:43.514926', false),
        ('ab5e707a-6ddf-44c6-9579-1d0746469c03', 'Abstract cumulative case rate by specimen date', 'Rate per 100,000 people of the total number of cases (people who have had at least one positive COVID-19 test result) since the start of the pandemic. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:18:25.107974', false),
        ('7d17fbee-557b-465d-bfbc-ff96b3ee2793', 'Abstract new cases PCR only rolling rate by specimen date', 'Rate per 100,000 people of the number of new cases (people who have had at least one positive COVID-19 test result) that were identified by a positive polymerase chain reaction (PCR) test, excluding people who had a positive rapid lateral flow (LFD) test within 3 days before the positive PCR test. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:35:08.126148', false),
        ('174e2372-b9dc-4b95-8a9e-068e0360759c', 'Abstract cumulative deaths within 28 days of a positive test rate by death date', 'Rate per 100,000 people of the total number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test since the start of the pandemic. Data are shown by the dates the deaths occurred.', '2021-09-03 13:37:20.758436', false),
        ('f0869ee4-ec96-4923-bdef-b515ef9a0d90', 'Abstract cumulative deaths within 28 days of a positive test by publish date', 'Total number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test since the start of the pandemic. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 13:39:13.838488', false),
        ('907b85a1-69a3-4709-a2ae-e56bd882f65c', 'Cumulative total number of deaths with COVID-19 on the death certificate, by area', 'Total number of deaths since the start of the pandemic of people whose death certificate mentioned COVID-19 as one of the causes, and death rates per 100,000 resident population.', '2021-06-23 10:17:06.724377', false),
        ('8005a2de-a1c8-4db5-aaf0-ec058a93bed6', 'Abstract cumulative deaths within 28 days of a positive test rate by publish date', 'Rate per 100,000 people of the total number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test since the start of the pandemic. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 13:41:11.961616', false),
        ('1921b765-217d-4ea5-80dd-a9077576f567', 'Abstract new virus tests', 'Daily numbers of new confirmed positive, negative or void COVID-19 virus test results. Tests are counted at the time they are processed.', '2021-09-03 15:02:56.641894', false),
        ('24de0103-3f5a-4f6f-b30d-3024fc93b6fe', 'Abstract new virus tests change percentage', 'The percentage change in the number of new confirmed positive, negative or void COVID-19 virus test results during the latest 7-day period, as a percentage of the number for the previous, non-overlapping 7-day period. Tests are counted at the time they are processed.', '2021-09-03 15:05:48.059135', false),
        ('5bd8241e-2846-4d17-89c2-360b2098b740', 'Abstract new virus tests direction', 'The direction of the change in the number of new confirmed positive, negative or void COVID-19 virus test results during the latest 7-day period compared to the previous, non-overlapping, 7-day period. Tests are counted at the time they are processed.

Positive changes mean numbers are increasing. These trends are shown with an upwards arrow. Negative changes mean numbers are decreasing. These trends are shown with a downwards arrow.', '2021-09-03 15:07:31.738060', false),
        ('27aaba54-9ff9-4c12-b483-5e37e0088704', 'Testing pillars', '# Testing pillars
The government''s mass testing programme provides testing through four routes known as pillars.

Information on the government’s national testing strategy, and the methodology used to report numbers of tests, is available on the Department for Health and Social Care website:

- [Coronavirus (COVID-19): scaling up testing programmes](https://www.gov.uk/government/publications/coronavirus-covid-19-scaling-up-testing-programmes)
- [Coronavirus (COVID-19): NHS Test and Trace statistics (England): methodology](https://www.gov.uk/government/publications/nhs-test-and-trace-statistics-england-methodology/nhs-test-and-trace-statistics-england-methodology)', '2021-09-20 14:11:49.831152', false),
        ('3e84456e-2640-4877-aab7-278e04ea384e', 'Daily number of deaths within 28 days of positive test and 7-day rates by age group (0-59 and 60+ years)', 'Daily number of deaths of people who had at least one positive test result for COVID-19 and died within 28 days of the first positive test, by date of death, and seven day rates of deaths per 100,000 population, grouped by age 0-59 and 60+ years.', '2021-06-23 10:14:03.080688', false),
        ('b720e947-d075-429f-bf60-a71eb274e419', 'R number and growth rate', 'The estimated R number (transmission rate) and daily infection growth rate.', '2021-06-24 09:36:19.978451', false),
        ('56a4679b-1a4b-49bd-8f43-28817a3f023d', 'Death rate by age group', 'Rate of deaths of people who had had a positive test result for COVID-19 and died within 28 days of the first positive test per 100,000 population by date of death, broken down by 5-year age group.', '2021-06-24 14:54:37.966437', false),
        ('43a03aa4-5570-4bf8-a18f-0389ab9bceab', 'PCR testing positivity', 'Polymerase chain reaction (PCR) tests are lab-based and test for the presence of SARS-CoV-2 virus. This data shows the number of people who received a PCR test in the previous 7 days, and the percentage of people who received a PCR test in the previous 7 days who had at least one positive PCR test result.

If a person has had more than one test result in the 7-day period, they are only counted once. If any of their tests in that period were positive, they count as one person with a positive test result. The positivity percentage is the number of people with a positive test result, divided by the number of people tested and multiplied by 100.', '2021-06-24 14:49:35.792106', false),
        ('6010e000-1ed4-4f0f-868f-5f12abde8f0e', 'Cumulative total number of COVID-19 associated deaths, by age and sex', 'Total number of deaths since the start of the pandemic of people who had had a positive test result for COVID-19, and death rates per 100,000 resident population. Some records have missing age or sex, so the sum of the subgroups does not equal the total deaths for the area.', '2021-06-23 10:13:12.572705', false),
        ('6670c5d8-d3de-4583-80ef-f355419c4c9d', 'Cumulative total number of cases, by area', 'Total number of people with a positive COVID-19 virus test result reported since the start of the pandemic, and rates of cases per 100,000 population.', '2021-06-23 09:39:24.303695', false),
        ('656e3eb8-25e7-43e7-987b-ec6441ccd3dc', 'Abstract cumulative cases LFD confirmed PCR by specimen date', 'Total number of cases (people who have had at least one positive COVID-19 test result) that were identified by a positive rapid lateral flow (LFD) test and confirmed by a positive polymerase chain reaction (PCR) test taken within 3 days. Data are shown by the date the lateral flow test was taken.', '2021-09-02 15:26:10.614274', false),
        ('91cbd380-b85f-4f62-9d60-c65cebab154a', 'Abstract new cases PCR only rolling sum by specimen date', 'The number of new cases (people who have had at least one positive COVID-19 test result) 
that were identified by a positive polymerase chain reaction (PCR) test, excluding people who had a positive rapid lateral flow (LFD) test within 3 days before the positive PCR test, within rolling 7-day periods. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:37:03.839393', false),
        ('257bf1ab-7c9d-4f8d-8c81-4a7713d37d52', 'Daily COVID-19 vaccinations given, by report date (metadata)', 'Data are reported daily, and include all vaccination events that are entered on the relevant system at the time of extract. Data are presented for vaccinations carried out up to and including the end of the report date.

The vaccination programme began on 8 December 2020 with people receiving the vaccine developed by Pfizer/BioNTech. People began receiving the Oxford University/AstraZeneca vaccine from 4 January 2021, and the Moderna vaccine from 7 April 2021. All 3 vaccines are given as 2 doses, at least 21 days apart, for a full vaccination course.

Initially the vaccines were prioritised to be administered to the over-80s, care home residents and workers, and NHS staff. The number of people of all ages who received each dose is reported.

# England
Vaccinations that were carried out in England are reported in the National Immunisation Management Service which is the system of record for the vaccination programme in England, including both hospital hubs and local vaccination services. Data are extracted at midnight on the date of report.

Following an IT issue reported to the NHS on 21 June, it was not possible to update vaccination figures for England. Vaccination figures reported for 22 June cover a 48-hour period.

# Northern Ireland
Data are extracted at the end of day of the date of report. Due to a processing issue, no data was reported for 22 March 2021. The newly reported number for 23 March 2021 includes vaccinations reported on 22 to 23 March 2021.

# Scotland
Vaccinations that were carried out in Scotland are reported in the Vaccination Management Tool. Data is extracted at 8:30am on the day following the date of report.

# Wales
Vaccinations that were carried out in Wales are reported in the Welsh Immunisation System. Data is extracted at 10pm on the date of report. No data was reported for 15 and 16 January 2021. The newly reported number for 17 January 2021 includes vaccinations reported on 15 to 17 January 2021.', '2021-06-30 11:40:39.073435', false),
        ('853eeb6e-8963-42fe-ae58-405314f43a6d', 'Abstract cumulative deaths within 60 days of a positive test by death date', 'Total number of deaths of people who had a positive test for COVID-19 and either died within 60 days of their first positive test or have COVID-19 mentioned on their death certificate since the start of the pandemic. Data are shown by the dates the deaths occurred.', '2021-09-03 13:44:12.549607', false),
        ('bd2223d8-c823-4176-85b3-7cc332d15f0b', 'Abstract cumulative deaths within 60 days of a positive test rate by death date', 'Rate per 100,000 people of the total number of deaths of people who had a positive test for COVID-19 and either died within 60 days of their first positive test or have COVID-19 mentioned on their death certificate since the start of the pandemic. Data are shown by the dates the deaths occurred.', '2021-09-03 13:46:44.896324', false),
        ('f9e8d4ce-6ede-4272-a854-c9a2f2a4f70c', 'Abstract new virus tests change', 'The difference between the number of new confirmed positive, negative or void COVID-19 virus test results during the latest 7-day period and the number for the previous, non-overlapping, 7-day period. Tests are counted at the time they are processed.', '2021-09-03 15:04:23.293254', false),
        ('361268ea-24a3-4fd3-b79f-6ad72abd706f', 'Abstract new virus tests rolling sum', 'The number of new confirmed positive, negative or void COVID-19 virus test results within rolling 7-day periods. Tests are counted at the time they are processed.', '2021-09-03 15:09:07.308357', false),
        ('0b8646ec-9b74-4098-bf22-7dd7b8a4b721', 'Cumulative total number of deaths within 28 days of positive test, by area', 'Total number of deaths since the start of the pandemic of people who had had a positive test result for COVID-19 and died within 28 days of the first positive test, and death rates per 100,000 resident population.', '2021-06-23 10:00:53.997305', false),
        ('f1891fb6-2a7d-4337-9b51-8174edac0aa3', 'Cases by age', 'Daily number of cases data aggregated by age into 0-59, 60 plus and individual five-year bands are available to download via the download data page. 7 day rolling averages and rates are also included where the data is presented by specimen date.', '2021-06-23 15:20:15.454050', false),
        ('a2138981-0341-4a38-9748-fb97c4f282ae', 'Number of vaccinations given, by report date', 'Number of COVID-19 vaccinations given, by report date.', '2021-06-24 09:22:13.855606', false),
        ('306c8e10-de1d-4517-bd0c-4c05f7fa7d5e', 'Cumulative total number of COVID-19 associated deaths, by area', 'Total number of deaths since the start of the pandemic of people who had had a positive test result for COVID-19, and death rates per 100,000 resident population.', '2021-06-23 10:14:49.402718', false),
        ('ca8cd65e-3270-40da-8a3c-a9bcd74de475', 'Daily deaths with COVID-19 on the death certificate, by date of death', 'Total number of deaths of people whose death certificate mentioned COVID-19 as one of the causes, by date of occurrence.', '2021-06-24 14:33:02.885101', false),
        ('b3a1ad60-f06c-4e23-ab07-4dcf6b607150', 'Vaccination reporting - old', 'Data are reported weekly on Thursdays, with data up to and including the previous Sunday. The vaccination programme began on 8 December 2020 with people receiving the vaccine developed by Pfizer/BioNTech, and people began receiving the Oxford University/AstraZeneca vaccine from 4 January 2021. Both vaccines are given as 2 doses, at least 21 days apart, for a full vaccination course.

Initially the vaccines were prioritised to be administered to the over-80s, care home residents and workers, and NHS staff. The number of people who received each dose is reported.

# UK
Data was provided by NHS England for the combined period 8-20 December 2020, which is included in the UK values for the week ending 20 December 2020. For this reason, no UK values are presented for week ending 13 December 2020. See the England section below for further details.

# England
Vaccinations that were carried out in England are reported in the National Immunisation Management Service which is the system of record for the vaccination programme in England, including both hospital hubs and local vaccination services. Data are extracted each Tuesday to reflect activity up to the close of the preceding Sunday.

At the start of the programme, NHS England could only provide data for the combined period 8-20 December. The implications are:

- No England, or UK figures are possible for week ending 13 December
- The England figure presented as week ending 20 December, includes the wider time period of 8-20 December
- The UK values presented as week ending 20 December, also include vaccinations in England from 10-13 December
- The cumulative totals by date are not affected

# Northern Ireland
As the vaccination programme began on 8 December 2020, the number of individuals reported to have been vaccinated in week ending 13 December 2020 only includes data from 8-13 December 2020.

# Scotland
Vaccinations that were carried out in Scotland are reported in the Vaccination Management Tool. As the vaccination programme began on 8 December 2020, the number of individuals reported to have been vaccinated in week ending 13 December 2020 only includes data from 8-13 December 2020.

# Wales
Vaccinations that were carried out in Wales are reported in the Welsh Immunisation System, and are extracted each Tuesday to reflect activity up to the close of the preceding Sunday. As the vaccination programme began on 8 December 2020, the number of individuals reported to have been vaccinated in week ending 13 December 2020 only includes data from 8-13 December 2020.', '2021-08-10 13:19:44.040224', false),
        ('434acb6d-855c-4136-afe2-97108ab0a2bb', 'Abstract cumulative admissions by age', 'Total number of patients admitted to hospital with COVID-19 since the start of the pandemic, by age.', '2021-09-01 08:49:47.135416', false),
        ('11bfef1b-8cca-410c-9da9-14e05acf32ab', 'Abstract cumulative cases LFD only by specimen date', 'Total number of cases (people who have had at least one positive COVID-19 test result) that were identified by a positive rapid lateral flow (LFD) test and were not confirmed by a positive polymerase chain reaction (PCR) test within 3 days. Data are shown by the date the lateral flow test was taken.', '2021-09-02 15:28:14.104049', false),
        ('0583c432-fbec-47ed-a599-9bb858175a50', 'Abstract cumulative deaths within 60 days of a positive test by publish date', 'Total number of deaths of people who had a positive test for COVID-19 and either died within 60 days of their first positive test or have COVID-19 mentioned on their death certificate since the start of the pandemic. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 13:49:25.315723', false),
        ('5aa97f83-ccd2-43e8-9d2e-ca3738a4fe71', 'Abstract planned antibody capacity by publish date', 'Estimated daily capacity for antibody serology testing reported by laboratories. Data are shown by the date the figures appeared in published totals.', '2021-09-03 15:10:38.468243', false),
        ('7db2bd04-026a-4e2b-aca1-a778ceb55206', 'Abstract planned capacity by publish date', 'Estimated daily capacity for testing reported by laboratories across all pillars of the UK government’s mass testing programme. 
Testing capacity is an estimate from labs of how many lab-based tests they have capacity to carry out each day based on availability of staff and resources. Data are shown by the date the figures appeared in published totals.', '2021-09-03 15:11:29.075391', false),
        ('6d43d864-2868-424c-8546-7b44d69c3c2d', 'Daily COVID-19 vaccinations given, by vaccination date (metadata)', 'Data are reported daily, and can be updated for historical dates as vaccinations given are recorded on the relevant system. Therefore, data for recent dates may be incomplete.

The vaccination programme began on 8 December 2020 with people receiving the vaccine developed by Pfizer/BioNTech. People began receiving the Oxford University/AstraZeneca vaccine from 4 January 2021, and the Moderna vaccine from 7 April 2021. All 3 vaccines are given as 2 doses, at least 21 days apart, for a full vaccination course.

Initially the vaccines were prioritised to be administered to the over-80s, care home residents and workers, and NHS staff. The number of people who received each dose is reported.

# England
Vaccinations that were carried out in England are reported in the National Immunisation Management Service which is the system of record for the vaccination programme in England. Only people who have an NHS number and are currently alive are included.

These are provided for population surveillance purposes, and will differ from NHS England daily outputs, which provide operational data for the management of the vaccination programme.

# Scotland
Vaccinations that were carried out in Scotland are reported to Public Health Scotland through the Vaccination Management Tool and General Practice IT systems.', '2021-06-30 11:45:12.356652', false),
        ('c9a8e41e-744b-49a0-abf9-9f5c92c51a68', 'By publish date', 'Reported by the date they were first included in the published totals.', '2021-07-22 09:02:01.176185', false),
        ('a06dd385-9bba-4922-994a-a7f05b018aed', 'Abstract new admissions change', 'The difference between the number of patients admitted to hospital with COVID-19 during the latest 7-day period and the number for the previous, non-overlapping, 7-day period.', '2021-09-01 09:04:35.667080', false),
        ('0ed8ef7c-683b-4f65-af36-2123e6a82948', 'Abstract new cases percentage change by publish date', 'The percentage change in the number of new cases (people who have had at least one positive COVID-19 test result) during the latest 7-day period, as a percentage of the number for the previous, non-overlapping 7-day period. Data are shown by the date the figures appeared in the published totals.', '2021-09-02 15:14:49.705289', false),
        ('cbde537a-f2aa-4e08-8478-9aba8d135f3c', 'Abstract cumulative cases PCR only by specimen date', 'Total number of cases (people who have had at least one positive COVID-19 test result) that were identified by a positive polymerase chain reaction (PCR) test, excluding people who had a positive rapid lateral flow (LFD) test within 3 days before the positive PCR test. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:28:48.266354', false),
        ('48ef48ab-ab8c-40c4-a2f1-90051f39a006', 'Abstract cumulative deaths within 60 days of a positive test rate by publish date', 'Rate per 100,000 people of the total number of deaths of people who had a positive test for COVID-19 and either died within 60 days of their first positive test or have COVID-19 mentioned on their death certificate since the start of the pandemic. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 13:51:25.438754', false),
        ('be52aa97-67ab-43cd-aa60-f1bfd49c17ff', 'Abstract planned PCR capacity by publish date', 'Estimated daily capacity for polymerase chain reaction (PCR) testing reported by laboratories across all pillars of the UK government’s mass testing programme. Testing capacity is an estimate from labs of how many lab-based tests they have capacity to carry out each day based on availability of staff and resources. Data are shown by the date the figures appeared in published totals.', '2021-09-03 15:12:54.219452', false),
        ('0ab46fea-efd5-4968-9e18-71ccc5fc0364', 'Deaths within 28 days - intro', 'Deaths of people who had a positive test result for COVID-19 and died with 28 days of their first positive test.

People who died more than 28 days after their first positive test are not included, whether or not COVID-19 was the cause of death.

Data can be presented by when someone died (date of death) or when the death was reported (date reported):

- deaths by date of death - each death is assigned to the date that the person died, however long it took for the death to be reported. Previously reported data are therefore continually updated
- deaths by date reported - each death is assigned to the date when it was first included in the published totals. The specific 24-hour periods reported against each date vary by nation', '2021-08-19 14:25:07.182248', false),
        ('0944a346-7bba-498d-8df8-7525d4ecb580', 'Testing capacity, by pillar', 'Estimated number of lab-based tests that can be performed based on staff availability and other resources, by testing pillar.', '2021-07-08 08:56:49.312046', false),
        ('5525607f-fc79-45f7-aacb-35680c98b9c7', 'Abstract new admissions change percentage', 'The percentage change in the number of patients admitted to hospital with COVID-19 during the latest 7-day period, reported as a percentage of the number for the previous, non-overlapping, 7-day period, by date reported.', '2021-09-01 09:07:52.918522', false),
        ('32a96dcc-2452-4b37-a502-9b0343c705fc', 'Abstract new cases change by specimen date', 'The difference between the number of new cases (people who have had at least one positive COVID-19 test result) during the latest 7-day period and the number for the previous, non-overlapping, 7-day period. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:31:30.350350', false),
        ('5f7e19ce-5e98-4558-8f10-71ca790563df', 'Cumulative cases', 'Total number of people with a positive COVID-19 virus test result reported since the start of the pandemic.', '2021-07-08 08:21:21.877879', false),
        ('04985bc9-fd3c-410e-8b64-30278b425b6a', 'Abstract capacity pillar 4', 'Reported testing capacity for pillar 4 – COVID-19 testing for national surveillance. 
For pillar 4, the reported capacity on a given date is the same as the number of tests processed by pillar 4 on that day.', '2021-09-03 14:01:47.172820', false),
        ('3b99ec88-c6fb-45c1-9fe8-3d1ee705c825', 'Abstract capacity pillar 3', 'Projected testing capacity for pillar 3 - antibody serology testing to show if people have antibodies from past infection with COVID-19.
Testing capacity is an estimate from labs of how many lab-based tests they have capacity to carry out each day based on availability of staff and resources.', '2021-09-03 14:05:53.399850', false),
        ('e5a04ec2-9183-4cc4-9a0d-36dc88ea3fa9', 'Abstract cumulative people vaccinated complete by publish date', 'Total number of people that have received a complete course of COVID-19 vaccination (2 doses given at least 21 days apart) since the start of the pandemic. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 15:19:28.764071', false),
        ('215af763-a965-4975-ad37-8d53b4c73ddd', 'By specimen date', 'Reported by the date when the sample was taken from the person being tested. Data for previous dates are regularly updated because of varying times between samples being taken and results being processed.', '2021-07-08 08:22:42.335658', false),
        ('6ef2cd51-ef0d-44e8-8d04-ea83c9781e0b', 'Pillar 1 testing', '# Pillar 1: NHS and PHE Testing

Virus testing in Public Health England (PHE) labs and NHS hospitals for those with a clinical need, and health and care workers. Pillar 1 data for England is provided by the NHS and PHE, and data from the devolved administrations are provided by the Department of Health of Northern Ireland, the Scottish Government, and Public Health Wales. Public Health Wales provide combined pillar 1 and partial pillar 2 data where tests are processed in NHS Wales labs.', '2021-07-08 08:36:52.467246', false),
        ('356afc1b-6890-4abf-afce-ba635762f107', 'Pillar 2 testing', '# Pillar 2: UK Government testing programme 

Virus testing for the wider population, as set out in government guidance.
Pillar 2 uses Lighthouse laboratories and has partnership arrangements with public, private and academic sector laboratories. Pillar 2 data for the UK (excluding Wales tests processed in NHS labs in Wales) is collected by commercial partners.
Pillar 2 data includes results from rapid lateral flow testing, where these have been reported to the NHS Digital Platform. Rapid lateral flow tests test for the presence of SARS-CoV-2 virus and are swab tests that give results in less than an hour, without needing to go to a laboratory.', '2021-07-08 08:37:32.587888', false),
        ('0bfa2634-9d1f-46d9-8ed1-ccacd49459ed', 'Pillar 3 testing', '# Pillar 3: Antibody testing 

Antibody serology testing to show if people have antibodies from past infection with COVID-19. Pillar 3 data is provided for England by NHS England and Improvement (NHSEI).', '2021-07-08 08:40:19.812293', false),
        ('6f898202-a613-4056-b43c-4d0c3c190fc4', '7-day figures', 'To help identify trends or patterns in the data over time, we calculate 7-day figures. These are updated daily, but even out random variations and  systematic variations over the week, for example rates being consistently lower at weekends.', '2021-07-15 09:09:42.867011', false),
        ('d7685d32-2c81-4574-b11f-8f73414774b5', 'Pillar 3 testing capacity', '# Pillar 3 testing capacity

Projected current capacity to process serology tests to show if people have antibodies from past infection with COVID-19.', '2021-08-19 15:16:50.861774', false),
        ('881dc1b9-e0fb-4303-b689-aa64866ec157', 'Deaths old', 'Deaths of people who had a positive test result for COVID-19.', '2021-07-15 14:34:29.232946', false),
        ('6c510f78-8106-4280-816c-04abd1ec7d39', 'Weekly deaths', 'Deaths reported each week.', '2021-08-03 07:28:13.057335', false),
        ('30be8c34-eae1-4cd7-a610-b70801832f6d', 'New tests', 'Daily number of new COVID-19 tests.', '2021-08-19 14:49:11.180709', false),
        ('9519e727-35ae-449e-8ad6-6cbf8734d414', 'Testing capacity, by pillar (metadata)', 'Estimated number of lab-based tests that can be performed based on staff availability and other resources, by testing pillar.

Projected laboratory capacity is an estimate of the number of COVID-19 tests each lab can process each day based on the availability of staff, chemical reagents and other resources required. These estimates are made locally by the labs themselves, aggregated and published weekly by the Department of Health and Social Care.
Testing capacity data are available by test type:

- lab-based virus tests that test for the SARS-CoV-2 virus. This includes lab-based pillar 1 and 2 tests and virus tests undertaken in pillar 4.
- antibody serology tests that test for COVID-19 antibodies. This includes pillar 3 tests and antibody serology tests undertaken in pillar 4.

Testing capacity data are available for the UK only so data cannot be presented separately for England, Northern Ireland, Scotland and Wales.
For pillars 1, 2 and 3, operational issues at labs can result in labs performing below capacity. There is also volatility in demand across the week, with significantly higher test volumes during the week and lower test volumes during the weekend.  Samples are moved between labs where there is a mismatch, but this can be limited by:

- the life of the sample
- distances between labs
- operational differences between labs (such as equipment or digital systems)

Comparing the number of tests processed on a given day or week to available lab capacity can give some understanding of how the programme is operating. However, you should do this with caution because different approaches are used to forecast capacity in each pillar.', '2021-07-21 11:43:37.707223', false),
        ('a9c2d970-96d5-446e-a294-0478f4ce682a', 'Abstract capacity pillar 1', 'Projected testing capacity for pillar 1 – NHS and PHE COVID-19 testing. 
Testing capacity is an estimate from labs of how many lab-based tests they have capacity to carry out each day, based on availability of staff and resources.', '2021-09-03 14:03:12.663410', false),
        ('a6949c95-664a-4c8a-8572-75ea3119b5f7', 'Abstract new admissions direction', 'The direction of the change in the number of new patients admitted to hospital with COVID-19 during the latest 7-day period compared to the previous, non-overlapping, 7-day period. 

Positive changes mean numbers are increasing. These trends are shown with an upwards arrow. Negative changes mean numbers are decreasing. These trends are shown with a downwards arrow.', '2021-09-01 09:11:45.816923', false),
        ('5f1811c6-5247-4f4a-913d-7e7b3ec2a378', 'Abstract cumulative cases by date reported', 'Total number of cases (people who have had at least one positive COVID-19 test result) since the start of the pandemic. Data are shown by the date the figures appeared in the published totals.', '2021-09-02 15:12:23.281419', false),
        ('813937d5-9b5a-4b89-9f54-ae96147e2b63', 'Pillar 4 testing', '# Pillar 4: Surveillance testing 

Virus testing and antibody serology testing for national surveillance supported by PHE, ONS, Biobank, universities and other partners to learn more about the prevalence and spread of the virus and for other testing research purposes, for example on the accuracy and ease of use of home testing. Pillar 4 data is collected by the NHS, PHE, and individual research study leads for the UK.', '2021-07-08 08:55:25.315296', false),
        ('d2fff977-fc0d-407a-8e6c-b90a6d09a4b6', 'Deaths death certificate', 'Deaths of people whose death certificate mentioned COVID-19 as one of the causes.', '2021-07-14 14:14:53.720178', false),
        ('f2fc177e-93e7-4b19-9821-1dc0877cc12c', 'By date reported', 'Reported by the date the death was first included in the published totals. 

The specific 24 hour periods reported against each date vary by nation and are given in the definition for ''Deaths within 28 days''.', '2021-07-14 14:18:00.077953', false),
        ('40fffe4f-1d0a-4490-bbf8-427fcf147a41', 'By date of registration', 'Reported by the date the death was registered.', '2021-07-14 14:18:53.874862', false),
        ('ce5a447b-9854-48fd-a82a-ee02cd4f946a', 'Rate definition', 'Rate per 100,000 population.', '2021-07-22 09:28:07.150458', false),
        ('659981b1-5d10-4d7e-a4a1-657af7e85b97', 'Test types', '# Test types

There are several different types of COVID-19 test. These include:

- polymerase chain reaction (PCR) tests
- rapid lateral flow (LFD) tests
- loop-mediated isothermal amplification (LAMP) tests

These tests differ in their sensitivity (the proportion of positives they correctly identify) and specificity (the proportion of negatives they correctly identify).

In certain situations and times of low prevalence, a confirmation PCR test is recommended to people with a positive rapid lateral flow test result. Self-reporting rapid lateral flow tests and cross-channel hauliers are always recommended to get a confirmation PCR, regardless of prevalence.

Test types are shown as:

- PCR (including LAMP and LamPORE tests)
- LFD with positive PCR (positive PCR result has been returned with a specimen date within 3 days of the positive rapid lateral flow test)
- LFD without positive PCR (positive PCR result has not been returned with a specimen date within 3 days of the positive rapid lateral flow test)', '2021-07-22 10:07:11.177613', false),
        ('af6aac87-3862-475a-b972-1247e755c64e', 'Cumulative deaths', 'Total number of deaths since the start of the pandemic.', '2021-07-23 08:22:19.837410', false),
        ('a7c6d0a5-1a64-447f-881f-21194444ef92', 'By date of death', 'Reported by the date the death occurred. 

Each death is assigned to the date that the person died, however long it took for the death to be reported. 

Previously reported data are therefore continually updated.', '2021-08-03 07:49:04.642588', false),
        ('b335041f-73d9-4b1a-a916-64df114f9311', 'New deaths', 'Daily number of new deaths.', '2021-08-19 14:47:29.103184', false),
        ('a3f58db1-964a-4dc4-9c30-f68b4091a19a', 'Cumulative tests', 'Total number of tests processed since the start of the pandemic.', '2021-07-21 12:59:23.590788', false),
        ('a5ed50b2-d1ed-4500-87f8-3615da318290', 'Abstract capacity pillar 1 and 2', 'Projected testing capacity for pillar 1 – NHS and PHE COVID-19 testing, and pillar 2 – COVID-19 testing under the UK Government testing programme.
Testing capacity is an estimate from labs of how many lab-based tests they have capacity to carry out each day based on availability of staff and resources.', '2021-09-03 14:04:30.893161', false),
        ('b1cfd094-b733-48dd-8a03-58ff04e29ebc', 'Abstract new cases direction by date reported', 'The direction of the change in the number of new cases (people who have had at least one positive COVID-19 test result) during the latest 7-day period compared to the previous, non-overlapping, 7-day period. Data are shown by the date the figures appeared in the published totals. 

Positive changes mean numbers are increasing. These trends are shown with an upwards arrow. Negative changes mean numbers are decreasing. These trends are shown with a downwards arrow.', '2021-09-02 15:15:34.349148', false),
        ('3f997a7c-079e-4062-8666-80e15e8ad128', 'Abstract cumulative people vaccinated complete by vaccination date', 'Total number of people that have received a complete course of COVID-19 vaccination (2 doses given at least 21 days apart) since the start of the pandemic. Data are shown by the date the second dose vaccination was given.', '2021-09-03 15:21:54.859282', false),
        ('298d2742-93ef-47fb-9a49-d9564f1ca2d0', 'Abstract new admissions rolling rate', 'The number of new patients admitted to hospital with COVID-19 within rolling 7-day periods, reported as a rate per 100,000 people.', '2021-09-01 09:20:36.739351', false),
        ('932b4404-29f7-4bb4-a03f-2f8df8774f07', 'Abstract new cases rolling rate by date reported', 'Rate per 100,000 people of the number of new cases (people who have had at least one positive COVID-19 test result) within rolling 7-day periods. Data are shown by the date the figures appeared in the published totals.', '2021-09-02 15:16:26.143674', false),
        ('0fa2913a-b2cd-4c1c-9853-efc7bbb3dbb4', 'Abstract capacity pillar 2', 'Projected testing capacity for pillar 2 – COVID-19 testing under the UK Government testing programme.
Testing capacity is an estimate from labs of how many lab-based tests they have capacity to carry out each day based on availability of staff and resources.', '2021-09-03 14:08:33.312869', false),
        ('d36fd4cf-1c6f-42f8-8b06-8ab5c29eb70e', 'Abstract cumulative antibody tests by publish date', 'Total number of confirmed positive, negative or void COVID-19 antibody test results since the start of the pandemic. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:09:19.555882', false),
        ('836a9b6c-edae-465a-ab56-fba221d7ce3d', 'Abstract cumulative LFD tests', 'Total number of confirmed positive, negative or void rapid lateral flow (LFD) test results for COVID-19 reported since the start of the pandemic.
Lateral flow tests give results without needing to go to a laboratory, so results are reported online by people who have taken the tests. Because of this, not all test results may be reported.', '2021-09-03 14:14:12.208352', false),
        ('1c960d8e-7777-427c-b800-e3103b99d760', 'Cases definition', '# Cases definition

COVID-19 cases are identified by taking specimens from people and testing them for the SARS-CoV-2 virus. If the test is positive this is referred to as a case. Some positive rapid lateral flow test results are confirmed with lab-based polymerase chain reaction (PCR) tests taken within 72 hours. If the PCR test results are negative, these are no longer reported as confirmed cases.
If a person has more than one positive test, they are only counted as one case.
Cases data includes all positive lab-confirmed polymerase chain reaction (PCR) test results plus, in England, positive rapid lateral flow tests that are not followed by a negative PCR test taken within 72 hours.

Figures may be shown by:

- date reported – the date it was first included in the published totals
	
- specimen date – the date the sample was taken from a patient

## UK

UK data include results from both pillar 1 and pillar 2 testing. See Testing Pillars, Pillar 1 Testing and Pillar 2 Testing for definitions.

## England

A positive case is someone with at least one confirmed positive test from a polymerase chain reaction (PCR) test, rapid lateral flow test or loop-mediated isothermal amplification (LAMP) test. Positive rapid lateral flow test results can be confirmed with PCR tests taken within 72 hours. If the PCR test results are negative, these are not reported as cases.

## Northern Ireland

A positive case is someone who has received a positive polymerase chain reaction (PCR) test result. If someone tests positive via a rapid lateral flow test, they must take a confirmatory PCR test. Positive rapid lateral flow tests are therefore not included in the figures for positive cases for Northern Ireland.

## Scotland

A positive case is someone with at least one confirmed positive polymerase chain reaction (PCR) test result. The number of rapid lateral flow test positive results are not included in daily case counts.

## Wales

A positive case is someone who has received a positive polymerase chain reaction (PCR) test result. If someone tests positive via a rapid lateral flow test, they are advised to take a confirmatory PCR test, so positive rapid lateral flow tests are not included in the figures for positive cases for Wales.

## Further information about the processes for counting cases in the devolved administrations

Details of the processes for counting cases in the devolved administrations are available on their websites:
- [Scottish Government coronavirus information](https://www.gov.scot/coronavirus-covid-19/)
- [Public Health Wales coronavirus information](https://public.tableau.com/profile/public.health.wales.health.protection#!/vizhome/RapidCOVID-19virology-Public/Headlinesummary)
- [Northern Ireland Department of Health coronavirus information](https://app.powerbi.com/view?r=eyJrIjoiZGYxNjYzNmUtOTlmZS00ODAxLWE1YTEtMjA0NjZhMzlmN2JmIiwidCI6IjljOWEzMGRlLWQ4ZDctNGFhNC05NjAwLTRiZTc2MjVmZjZjNSIsImMiOjh9)', '2021-07-22 09:08:48.958821', false),
        ('731dbc5b-fb93-4321-a7bc-0c906cce974b', 'Deaths 60 days', 'Deaths of people who had a positive test result for COVID-19 and either died within 60 days of the first positive test or have COVID-19 mentioned on their death certificate.

This metric is included in the dashboard for comparison to the headline 28-day metric.

Full [details of the methodology are available in the technical summary of the PHE data series on deaths in people with COVID-19](https://www.gov.uk/government/publications/phe-data-series-on-deaths-in-people-with-covid-19-technical-summary)', '2021-08-03 07:51:37.441823', false),
        ('58593e53-efd1-4c80-bcc8-7fa83262ca38', 'Deaths ONS - metadata', '# Office for National Statistics (ONS) deaths

Data for England and Wales are published weekly by the ONS. Deaths registered up to Friday are published 11 days later on the Tuesday.

Deaths are allocated to the person''s usual area of residence.

Bank holidays can affect the number of deaths registered in a given week.

Details of collection and further data are available on the ONS website: [Office for National Statistics](https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/conditionsanddiseases/datalist?filter=datasets)', '2021-08-19 14:21:45.755346', false),
        ('d626ef38-4a5b-4acd-8833-33a6e370f036', 'Abstract cumulative people vaccinated 1st dose by publish date', 'Total number of people that have received a first dose COVID-19 vaccination since the start of the pandemic. A complete COVID-19 vaccination course is 2 doses given at least 21 days apart. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 15:24:24.696223', false),
        ('12180df5-447e-4abe-875b-5f8c011cfd33', 'Chart abstract for test types', 'Number of people with at least one positive COVID-19 test result. Cases are shown by type of test used in their first positive test and by specimen date.
      
Positive rapid lateral flow test results can be confirmed with PCR tests taken within 72 hours. If the PCR test results are negative, these are not reported as cases.
      
The test types shown are: 

- lab-based polymerase chain reaction (PCR)
- rapid lateral flow test (LFD) with positive PCR (this means the result has been verified
      with a positive PCR result taken within 3 days)
- LFD only (no PCR taken within 3 days)

      
People tested positive more than once are only counted once, on the date of their first positive test. 

Data for the period ending 5 days before the date when the website was last updated with data for the selected area are incomplete. Some LFD results may change or be removed as more PCR results are reported.', '2021-08-31 12:11:05.590538', false),
        ('a584ba3d-6dcd-4f98-a844-42aaca72939f', 'Abstract cumulative admissions', 'Total number of patients admitted to hospital with COVID-19 since the start of the pandemic.', '2021-09-01 08:46:52.032232', false),
        ('b29e0208-6d27-4e67-9ae5-37a2505ccfaa', 'Abstract new admissions rolling sum', 'The number of new patients admitted to hospital with COVID-19 within rolling 7-day periods.', '2021-09-01 09:23:35.645272', false),
        ('f2852a48-0024-40ea-94c1-d78519497a03', 'Abstract new cases rolling sum by date reported', 'The number of new cases (people who have had at least one positive COVID-19 test result) within rolling 7-day periods. Data are shown by the date the figures appeared in the published totals.', '2021-09-02 15:17:09.975350', false),
        ('5226fbc9-b96e-423c-9944-d04355a611dc', 'Abstract new cases by specimen date', 'Daily numbers of new cases (people who have had at least one positive COVID-19 test result). Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:19:00.827820', false),
        ('aef3f60b-4626-4491-b4eb-a23bca4a9b83', 'Abstract new cases age demographics by specimen date', 'Age and sex breakdown of the daily numbers of new cases (people who have had at least one positive COVID-19 test result). Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:19:36.877174', false),
        ('f569d3d1-5956-4e84-99d9-650e35c9d308', 'Abstract cumulative PCR tests by publish date', 'Total number of confirmed positive, negative or void polymerase chain reaction (PCR) test results for COVID-19 since the start of the pandemic.', '2021-09-03 14:13:13.553272', false),
        ('5da8d33b-9a4e-4ca6-b416-432dbf043c84', 'Geographic allocation of cases and deaths', '## Geographic allocation of cases and deaths
Cases and deaths can be removed or reallocated as records are regularly updated with new information, which can create apparent discrepancies. If numbers for an area are low, these changes can result in negative numbers, which are shown as 0. However, the changes would be included for areas where a negative number would not be created.

For example: Local authority A reports 2 new cases. Local authority B reports no new cases, but 1 previously reported case is reallocated to another local authority. This shows as 0 new cases reported for local authority B.
These local authorities form a larger area. 1 case is shown in this larger area (2 cases in local authority A, minus 1 case in local authority B).

PHE has published a [comparison of geographic allocation methodologies](https://www.gov.uk/government/publications/covid-19-comparison-of-geographic-allocation-of-cases-in-england-by-lower-tier-local-authority) for cases by specimen date at lower tier local authority level.

Due to their small populations, counts for City of London and Isles of Scilly are combined with Hackney and Cornwall respectively when presented at local authority level, in order to prevent disclosure control issues.', '2021-09-20 14:10:45.136469', false),
        ('3ba669ea-1b26-4c31-ac6f-447f362d6240', 'Age and sex', 'By age and sex. 

Some records have missing age or sex, so the sum of the subgroups does not equal the total deaths for the area.', '2021-07-23 12:20:08.094504', false),
        ('bc5109ea-4aa6-4710-85d0-bceb0374baad', 'Abstract for male and female cases', 'Number of people with at least one positive
COVID-19 test result, either lab-reported or lateral flow device (England only) since the
beginning of the pandemic, by age and sex. 

Positive rapid lateral flow test results can be confirmed with PCR tests taken within 72 hours. If the PCR test results are negative, these are not reported as cases.', '2021-07-26 14:11:41.795270', false),
        ('9c570352-cc53-43fb-8e52-17feb1ef79de', 'New people vaccinated', 'Number of people newly vaccinated, reported each day.', '2021-08-19 15:40:20.366895', false),
        ('de7a0904-9058-410b-b362-a7b320409cc8', 'Cumulative people vaccinated', 'Total number of people vaccinated since the start of the pandemic.', '2021-08-10 10:17:45.380363', false),
        ('10aa6845-567d-4c93-9c28-99148194a60b', 'First dose', 'The number of people who have received a first dose COVID-19 vaccination.', '2021-08-10 10:18:58.668758', false),
        ('fb3b9f79-5047-42c2-b07a-1be2c581e28d', 'Second dose', 'The number of people who have received a second dose COVID-19 vaccination.', '2021-08-10 10:19:22.409745', false),
        ('2671bffc-6ad8-4dc0-85cb-163b14f3af5e', 'Complete vaccination', 'The number of people who have received a complete course of COVID-19 vaccination.

A complete COVID-19 vaccination course is 2 doses given at least 21 days apart.', '2021-08-10 10:19:52.797821', false),
        ('6f9ed5aa-c29e-44ec-ae7c-618888d2f411', 'Cumulative vaccines given', 'Total number of COVID-19 vaccinations given (both first and second dose) since the start of the pandemic.', '2021-08-10 13:33:23.086997', false),
        ('5e170b46-83ed-4b90-a6a5-432738f7f7ee', 'Abstract new admissions', 'Daily number of new admissions to hospital of patients with COVID-19.', '2021-09-01 09:01:36.397192', false),
        ('463a8ab0-58ac-487a-9aeb-abcd39c707e9', 'Abstract new cases by date reported', 'Daily numbers of new cases (people who have had at least one positive COVID-19 test result). Data are shown by the date the figures appeared in the published totals.', '2021-09-02 15:13:29.703038', false),
        ('b5925703-0581-4e6f-9447-f1ebf194c273', 'Vaccination uptake', 'Number of people and percentage of the population aged 16 and over who have received a COVID-19 vaccination.', '2021-08-19 15:45:51.756535', false),
        ('300974ce-5285-4f13-885e-0e8f786ae211', 'Abstract new cases percentage change by specimen date', 'The percentage change in the number of new cases (people who have had at least one positive COVID-19 test result) during the latest 7-day period, as a percentage of the number for the previous, non-overlapping, 7-day period. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:20:22.988966', false),
        ('cd90ffb7-dece-43e7-82a1-3270f1e48dfa', 'Abstract new deaths within 28 days of a positive test by death date', 'Daily numbers of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test. Data are shown by the dates the deaths occurred.', '2021-09-03 12:39:44.080389', false),
        ('bf921477-f663-40df-ad35-39822daaab03', 'Abstract cumulative pillar 4 tests by publish date', 'Total number of confirmed positive, negative or void COVID-19 tests conducted under pillar 4 – COVID-19 testing for national surveillance, since the start of the pandemic. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:15:37.224603', false),
        ('cfa2b226-f636-4d22-98dd-fdd8729bdefb', 'Abstract cumulative pillar 1 tests by publish date', 'Total number of confirmed positive, negative or void COVID-19 tests conducted under pillar 1 – NHS and PHE COVID-19 testing, since the start of the pandemic. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:17:20.076320', false),
        ('b6b1a96e-afda-4f2d-b9f5-00bcaea4afd3', 'Deaths NSO - intro', 'Provisional counts of the number of deaths registered in the UK where COVID-19 is mentioned as a cause on the death certificate. The deceased may not have had a confirmed positive test for COVID-19. People who had had COVID-19 but did not have it mentioned on their death certificate as a cause of death are excluded.', '2021-08-19 14:11:52.013824', false),
        ('8f97317b-4e5c-4a2e-9bc2-f3c73d5951fe', 'Pillar 2 testing capacity', '# Pillar 2 testing capacity
## Lighthouse laboratories 
Capacity is the maximum number of samples the Lighthouse laboratories estimate they will be able to process the following day. The figure takes into account equipment, laboratory facilities, anticipated machine down time, and staffing.

## Partner laboratories 
Capacity for partner laboratories is the daily number of tests that the laboratory has been contracted to complete.', '2021-08-19 15:16:09.177404', false),
        ('22bd2267-5332-4b99-9437-91370f7ebf8f', 'Pillar 4 testing capacity', '# Pillar 4 testing capacity

The reported pillar 4 capacity on a given date is the same as the number of tests processed by pillar 4 on that day.', '2021-08-19 15:17:24.047854', false),
        ('3a4d7aee-8fa4-4554-8809-70803de32370', 'Vaccination uptake (metadata)', '# UK and nations vaccination uptake - metadata
Uptake percentages for the UK and nations are shown by report date. Percentage uptake by report date is calculated by dividing the total number of vaccinations given to people of all ages by the mid-year 2020 population estimate for people aged 16 and over, published by the Office for National Statistics.

The percentage uptake published here for the nations of the UK may differ from those reported by the nations individually. In particular, figures published by Public Health Wales use denominators of people registered with NHS Wales rather than mid-year population estimates.

Percentage uptake for age groups in England is calculated on a different basis, and presented for the purposes of population surveillance and comparison with regional and local authority data.

## England local areas and age breakdowns
Uptake percentages for regions and local authorities, along with age breakdowns for England and local areas within England, are shown by vaccination date. Percentage uptake by vaccination date is calculated by dividing the total number of vaccinations given to people aged 16 and over by the number of people aged 16 and over on the National Immunisation Management Service (NIMS).

Vaccinations carried out in England are reported in NIMS, the system of record for the vaccination programme in England. Only people who have an NHS number and are currently alive are included.

These are provided for population surveillance purposes, and will differ from NHS England daily outputs, which provide operational data for the management of the vaccination programme.

## Scotland local areas
Uptake percentages for local authorities are presented by vaccination date. Percentage uptake by vaccination date is calculated by dividing the total number of vaccinations given to people aged 16 and over by the mid-year 2020 population estimate for people aged 16 and over.', '2021-08-19 15:50:14.344357', false),
        ('5d5b9775-c041-4828-a47a-4841dd0bce0c', 'Vaccination reporting', '# Vaccination reporting
Data can be presented by:

- when someone was vaccinated (vaccination date) – each vaccination is assigned to the date it was given, however long it took for the vaccination to be reported. Previously reported data are therefore continually updated

- when the vaccination was reported (date reported) - each vaccination is assigned to the date when it was first included in the published totals

Daily figures include all vaccines that were given up to and including the date shown, and that were entered on the relevant system at the time of extract.

Data are reported weekly on Thursdays, with data up to and including the previous Sunday.  

The vaccination programme began on 8 December 2020.

All 3 vaccines are given as 2 doses, at least 21 days apart, for a full vaccination course.

As the vaccination programme began on 8 December 2020, the number of individuals reported to have been vaccinated in week ending 13 December 2020 only includes data from 8-13 December 2020.

The number of people who received each dose is reported.

## UK
Data was provided by NHS England for the combined period 8 to 20 December 2020, which is included in the UK values for the week ending 20 December 2020. For this reason, no UK values are shown for week ending 13 December 2020. See the England section for further details.

## England
Vaccinations carried out in England are reported in the National Immunisation Management Service. This is the system of record for the vaccination programme in England, including both hospital hubs and local vaccination services. Data are published daily and include vaccinations administered up to midnight on the previous day. 

At the start of the programme, NHS England could only provide data for the combined period 8 to 20 December. Because of this:

- no England or UK figures are possible for week ending 13 December
- the England figure presented as week ending 20 December, includes the wider time period of 8 to 20 December
- the UK values presented as week ending 20 December, also include vaccinations in England from 10 to 13 December
- the cumulative totals by date are not affected

## Northern Ireland
Northern Ireland data are derived from the Northern Ireland Vaccinations Management System. Data are reported daily. 

## Scotland
Vaccinations carried out in Scotland are recorded in the Vaccination Management Tool. Data are reported daily and include vaccinations administered up to 8:30am on the date reported. 

## Wales
Vaccinations carried out in Wales are reported in the Welsh Immunisation System. Data are reported daily and include vaccinations recorded up to 10pm on the previous day.', '2021-08-19 15:54:48.049722', false),
        ('4436563d-a0ab-4ebc-a0c0-2b4edaae8ce9', 'By vaccination date', 'Reported by the date the vaccination was given.', '2021-08-10 10:27:58.931602', false),
        ('e6f67fd0-9bde-4b06-811f-54293dfb8477', 'Abstract new cases direction by specimen date', 'The direction of the change in the number of new cases (people who have had at least one positive COVID-19 test result) during the latest 7-day period compared to the previous, non-overlapping, 7-day period. Data are shown by the date the sample was taken from the person being tested.

Positive changes mean numbers are increasing. These trends are shown with an upwards arrow. Negative changes mean numbers are decreasing. These trends are shown with a downwards arrow.', '2021-09-02 15:21:46.929663', false),
        ('a0ca6308-1def-433b-83b0-8a8d53cec2ef', 'Abstract new deaths within 28 days age demographics by death date', 'Daily numbers of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test, broken down by age and sex. Data are shown by the dates the deaths occurred.', '2021-09-03 12:44:05.290963', false),
        ('f3aa03d4-7076-49a3-939b-d7fc829efeb7', 'Abstract cumulative pillars 1 and 2 tests by publish date', 'Total number of confirmed positive, negative or void COVID-19 tests conducted under pillar 1 – NHS and PHE COVID-19 testing and pillar 2 – the UK Government COVID-19 testing programme, since the start of the pandemic. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:22:58.257468', false),
        ('f5a8faa8-313e-4121-8722-c7aaaf84793c', 'Abstract cumulative pillar 3 tests by publish date', 'Total number of confirmed positive, negative or void COVID-19 antibody tests conducted under pillar 3 – antibody serology testing, since the start of the pandemic. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:24:36.377624', false),
        ('c669fbec-c8d9-4190-8363-60ce3cd7968d', 'Hospital admissions', '# Hospital admissions

Data are not reported by each nation every day. This data is classified as management information rather than official statistics.
Use caution when interpreting the data as there are inconsistencies between England, Northern Ireland, Scotland and Wales:

-	definitions are not always consistent
-	for England, no revisions have been made to the dataset. When known errors come to light, Trusts make the appropriate correction in the next day’s data. For Wales, Northern Ireland and Scotland, historical data are subject to revision

## England
England data include people admitted to hospital who tested positive for COVID-19:

- in the 14 days before their admission
- during their stay in hospital
Inpatients diagnosed with COVID-19 after admission are reported as being admitted on the day before their diagnosis.

Admissions figures include people admitted to:

-	NHS acute hospitals
-	mental health and learning disability trusts
-	independent service providers commissioned by the NHS
Data are reported daily by trusts to NHS England and NHS Improvement. 
You can find full NHS definitions in the estimated admissions section of the Publication Definitions document available on the [NHS COVID-19 Hospital Activity website](https://www.england.nhs.uk/statistics/statistical-work-areas/covid-19-hospital-activity/). Reporting dates reflect admissions and new inpatient diagnoses for the previous day. 
Daily data at NHS Trust level is reported weekly on Thursdays.

## Northern Ireland
Northern Ireland data include confirmed COVID-19 hospital admissions by admission date. If an inpatient tests positive for COVID-19, their admission code is revised to count them as a confirmed COVID-19 admission. 
Data are routinely uploaded to the Regional Data Warehouse from the Patient Administrative System (PAS), a patient-level administrative data source from health and social care hospitals in Northern Ireland. Data are sourced from the Regional Data Warehouse by the Department of Health Northern Ireland and published daily Monday to Friday. 

## Scotland
Data for Scotland include hospital admissions for patients who tested positive for COVID-19:

-	in the 14 days before their admission
-	on the day of their admission
-	during their stay in hospital
Data are collected from all Scottish NHS Health Boards and Golden Jubilee Hospital National Facility by the RAPID (Rapid Preliminary Inpatient Data) dataset and are usually reported daily Monday to Saturday.    

## Wales
Wales data include confirmed and suspected cases. Figures include admissions to hospital in the previous 24-hour period up to 9am.
The numbers of admissions are not comparable with other nations.
Data are collected in a daily return from Digital Health and Care Wales (DHCW) and are reported daily Monday to Friday by StatsWales.', '2021-09-20 14:24:09.295442', false),
        ('d8e904bb-a2d7-454e-8b5e-6d455b88c011', 'Weekly people vaccinated', 'Number of new people vaccinated each week.', '2021-08-10 12:25:36.555406', false),
        ('37620794-9433-4988-992f-27559f8d5909', 'Vaccination uptake (metadata - old)', 'Data are reported daily, and include all vaccination events that are entered on the relevant system at the time of extract.

# UK and nations headline uptake
Headline uptake percentages for the UK and nations are presented by report date. Percentage uptake by report date is calculated by dividing the total number of vaccinations given to people of all ages, by the mid-year 2019 population estimate for people aged 18 and over, published by the Office for National Statistics.

The percentage uptake published here for the nations of the UK may differ from those reported by the nations individually. In particular, figures published by Public Health Wales use denominators of those registered with NHS Wales rather than mid-year population estimates.

Percentage uptake for age groups in England is calculated on a different basis, and presented for the purposes of population surveillance and comparison with regional and local authority data (see below).

# England local areas and age breakdowns
Uptake percentages for regions and local authorities, along with age breakdowns for England and local areas within England, are presented by vaccination date. Percentage uptake by vaccination date is calculated by dividing the total number of vaccinations given to people aged 18 and over by the number of people aged 18 and over on the National Immunisation Management Service (NIMS).

Vaccinations that were carried out in England are reported in NIMS which is the system of record for the vaccination programme in England. Only people who have an NHS number and are currently alive are included.

These are provided for population surveillance purposes, and will differ from NHS England daily outputs, which provide operational data for the management of the vaccination programme.

# Scotland local areas
Uptake percentages for local authorities are presented by vaccination date. Percentage uptake by vaccination date is calculated by dividing the total number of vaccinations given to people aged 18 and over by the mid-year 2019 population estimate for people aged 18 and over.', '2021-08-10 13:19:06.687576', false),
        ('8db4bffe-e0e5-4f80-b2af-37134817a452', 'New vaccines given', 'Number of new COVID-19 vaccinations given (both first and second dose) each day.', '2021-08-10 13:36:08.094992', false),
        ('c47cff32-ae34-4c5a-9327-f8c7a11351ba', 'Deaths within 28 days - metadata', '# Deaths within 28 days of a positive COVID-19 test – metadata 
Deaths are allocated to the person''s usual area of residence.

## England

Data on COVID-19 associated deaths in England are produced by Public Health England (PHE) from multiple sources linked to confirmed case data. Deaths newly reported each day cover the 24 hours up to 5pm on the previous day.

Deaths are only included if the deceased had a positive test for COVID-19 and died within 28 days of the first positive test.

Regional, upper tier local authority (UTLA) and lower tier local authority (LTLA) death counts exclude England deaths for which the exact location of residence is unknown. Therefore, they may not add up to the England total.

Postcode of residence for deaths is collected at the time of testing. This is supplemented, where available, with information from ONS mortality records, Health Protection Team reports and NHS Digital Patient Demographic Service records.

Full details of the methodology are available in the [technical summary of the PHE data series on deaths in people with COVID-19](https://www.gov.uk/government/publications/phe-data-series-on-deaths-in-people-with-covid-19-technical-summary)

## Northern Ireland
Data for Northern Ireland include all cases reported to the Public Health Agency (PHA) where the deceased had a positive test for COVID-19 and died within 28 days. PHA sources include reports by healthcare workers (eg Health and Social Care Trusts, GPs) and information from local laboratory reports. Deaths reported against each date cover the 24 hours up to 9:30am on the same day.

## Scotland
Data for Scotland include deaths in all settings which have been registered with National Records of Scotland (NRS) where a laboratory-confirmed report of COVID-19 in the 28 days prior to death exists. Deaths reported against each date cover the 24 hours up to 9:30am on the same day.

## Wales
Data for Wales include reports to Public Health Wales of deaths of hospitalised patients in Welsh Hospitals or care home residents where COVID-19 has been confirmed with a positive laboratory test and the clinician suspects this was a causative factor in the death. The figures do not include:

- people who may have died from COVID-19 but COVID-19 was not confirmed by laboratory testing
- people who died outside of hospital or care home settings
- Welsh residents who died outside of Wales
Deaths reported each day cover the 24 hours up to 5pm on the previous day. The majority of deaths included occur within 28 days of a positive test result.', '2021-08-19 14:31:26.445502', false),
        ('ca6448b1-1755-4ffb-8e23-af1ab336be35', 'Abstract new cases rolling rate by specimen date', 'Rate per 100,000 people of the number of new cases (people who have had at least one positive COVID-19 test result) within rolling 7-day periods. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:22:26.705990', false),
        ('22bba168-3548-484e-9574-f4db9cae7d7d', 'Abstract new deaths within 28 days of a positive test rate by death date', 'Rate per 100,000 people of the daily numbers of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test. Data are shown by the dates the deaths occurred.', '2021-09-03 12:46:41.578325', false),
        ('81464889-7f52-4e53-b332-10d3e4c8a266', 'Abstract cumulative pillar 2 tests by publish date', 'Total number of confirmed positive, negative or void COVID-19 tests conducted under pillar 2 – the UK Government COVID-19 testing programme, since the start of the pandemic. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:25:54.638557', false),
        ('ba8e1a11-a4ca-4914-8cfd-2c34bd30a054', 'Abstract new cases rolling sum by specimen date', 'The number of new cases (people who have had at least one positive COVID-19 test result) within rolling 7-day periods. Data are shown by the date the sample was taken from the person being tested.', '2021-09-02 15:23:06.568851', false),
        ('a83f96b8-e713-4eb3-902b-1a3574d9dab9', 'Deaths NSO - metadata', '# Deaths data from the national statistics bodies - metadata

Data are published weekly by the national statistics bodies in England, Northern Ireland, Scotland and Wales. Deaths registered up to Friday are published 11 days later on the Tuesday.

Coding of deaths by cause for the latest week is incomplete. 

Deaths are allocated to the person''s usual area of residence.

Bank holidays can affect the number of deaths registered in a given week.

Details of collection and further data are available on the national statistics bodies'' websites:

## United Kingdom
England, Wales and Northern Ireland weekly deaths run from Saturday to Friday. Scotland deaths run from Monday to Sunday. This means deaths for individual days will not add up to the weekly total.

## England and Wales
[Office for National Statistics.](https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/conditionsanddiseases/datalist?filter=datasets)

The number of deaths by date of occurrence is based on deaths registered up to a later registration date to take into account registration delay. Read more about [registration delays on the Office for National Statistics website.](https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths/articles/impactofregistrationdelaysonmortalitystatisticsinenglandandwales/2019)

## Northern Ireland
[Northern Ireland Statistics and Research Agency](https://www.nisra.gov.uk/statistics/ni-summary-statistics/coronavirus-covid-19-statistics)

## Scotland
[National Records of Scotland](https://www.nrscotland.gov.uk/statistics-and-data/statistics/statistics-by-theme/vital-events/general-publications/weekly-and-monthly-data-on-births-and-deaths/deaths-involving-coronavirus-covid-19-in-scotland)', '2021-08-19 14:32:24.965104', false),
        ('6f483279-e356-4036-aae0-215b1270d101', 'Abstract new deaths within 28 days of a positive test rolling rate by death date', 'Rate per 100,000 people of the number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test within rolling 7-day periods. Data are shown by the dates the deaths occurred.', '2021-09-03 13:01:43.754964', false),
        ('d43a4ee6-e0a1-418e-a8d4-ed18b32dbf40', 'Abstract new deaths within 28 days of a positive test rolling sum by death date', 'The number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test within rolling 7-day periods. Data are shown by the dates the deaths occurred.', '2021-09-03 13:04:16.973711', false),
        ('484e011a-9dff-490a-a22d-6daa8af286a7', 'Abstract cumulative virus tests', 'Total number of confirmed positive, negative or void COVID-19 virus test results since the start of the pandemic. Tests are counted at the time they are processed.', '2021-09-03 14:30:35.686947', false),
        ('ba3a1832-d0b9-40b9-bbdf-62e09b97f70a', 'Abstract new LFD tests by publish date', 'Daily numbers of new confirmed positive, negative or void rapid lateral flow (LFD) test results for COVID-19. Lateral flow tests give results without needing to go to a laboratory, so results are reported online by people who have taken the tests. Because of this, not all test results may be reported. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:33:57.281660', false),
        ('6716d107-a7de-46a5-b853-62c429bf3806', 'Abstract new cases LFD confirmed PCR by specimen date', 'Daily number of new cases (people who have had at least one positive COVID-19 test result) that were identified by a positive rapid lateral flow (LFD) test and were confirmed by a positive polymerase chain reaction (PCR) test taken within 3 days. Data are shown by the date the lateral flow test was taken.', '2021-09-02 15:24:04.360760', false),
        ('45480184-f541-4332-abc2-b570c6490fd5', 'Abstract new deaths within 28 days of a positive test by publish date', 'Daily numbers of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 13:06:31.593645', false),
        ('8c14be0a-6282-4932-9968-f5e1cda14b4e', 'Abstract new deaths within 28 days of a positive test change by publish date', 'The difference between the number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test during the latest 7-day period and the number for the previous, non-overlapping, 7-day period. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 13:08:18.887348', false),
        ('d19c55fa-ba41-4b89-8618-68278208ac30', 'Abstract new antibody tests by publish date', 'Daily numbers of new confirmed positive, negative or void COVID-19 antibody test results. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:32:10.606416', false),
        ('3e5aeec4-55d4-4e70-a003-4e9901186300', 'Abstract new PCR tests by publish date', 'Daily numbers of new confirmed positive, negative or void polymerase chain reaction (PCR) test results. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:35:43.740841', false),
        ('6b8ca2a3-ea38-4438-bb88-4b1de230d318', 'Abstract new PCR tests change by publish date', 'The difference between the number of new confirmed positive, negative or void polymerase chain reaction (PCR) test results during the latest 7-day period and the number for the previous, non-overlapping, 7-day period. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 14:37:20.341573', false),
        ('57d209c8-b171-41f3-acd1-add0df96557a', 'Cumulative total number of deaths within 60 days of positive test, by age and sex', 'Total number of deaths since the start of the pandemic of people who had had a positive test result for COVID-19 and either died within 60 days of the first positive test or have COVID-19 mentioned on their death certificate, and equivalent death rates per 100,000 resident population. Some records have missing age or sex, so the sum of the subgroups does not equal the total deaths for the area.

This metric is included in the dashboard for comparison to the headline 28-day metric.

Full details of the methodology are available on [GOV.UK.](https://www.gov.uk/government/publications/phe-data-series-on-deaths-in-people-with-covid-19-technical-summary)', '2021-06-23 10:11:52.637956', false),
        ('be10845d-3089-47ae-b4b9-8ba5ee0f8c95', 'Daily COVID-19 vaccinations given, by report date', 'Number of people who have received a COVID-19 vaccination, by report date.', '2021-06-24 09:20:11.278921', false),
        ('459db480-a818-488c-9563-fca24ea5b16d', 'Cumulative total number of cases, by age and sex', 'Total number of people with a positive COVID-19 virus test result reported since the start of the pandemic, and rates of cases per 100,000 resident population. Some test results have missing age or sex, so the sum of the subgroups does not equal the total cases for the area.', '2021-06-23 09:37:35.350303', false),
        ('5c80164d-67e5-460d-938e-a4f260564e41', 'Vaccines given', 'The number of COVID-19 vaccinations given (both first and second dose).', '2021-08-19 15:56:27.185850', false),
        ('14760193-a1ef-4f32-9469-7214c5bcfab4', 'Abstract new cases LFD confirmed by PCR rolling rate by specimen date', 'Rate per 100,000 people of the number of new cases (people who have had at least one positive COVID-19 test result) that were identified by a positive rapid lateral flow (LFD) test and were confirmed by a positive polymerase chain reaction (PCR) test taken within 3 days, within rolling 7-day periods. Data are shown by the date the lateral flow test was taken.', '2021-09-02 15:24:38.511156', false),
        ('6f74511f-c1dc-4b9a-a23a-1c7e2f59331f', 'Abstract new cases LFD only by specimen date', 'Daily numbers of new cases (people who have had at least one positive COVID-19 test result) that were identified by a rapid lateral flow (LFD) test and were not confirmed by a positive polymerase chain reaction (PCR) test within 3 days. Data are shown by the date the lateral flow test was taken.', '2021-09-02 15:26:50.611517', false),
        ('1cd6cfcf-1762-4d8d-9856-e1ddfa4ff4df', 'Abstract new deaths within 28 days of a positive test change percentage by publish date', 'The percentage change in the number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test during the latest 7-day period, as a percentage of the number for the previous, non-overlapping, 7-day period. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 13:11:12.103105', false),
        ('4deceb54-8487-44e6-9e2c-98b830ba5193', 'Abstract new deaths within 60 days of a positive test by death date', 'Daily numbers of deaths of people who had a positive test for COVID-19 and either died within 60 days of their first positive test or have COVID-19 mentioned on their death certificate. Data are shown by the dates the deaths occurred.', '2021-09-03 13:17:57.973354', false),
        ('db4011b2-d465-40bb-871d-07b1eeb3e6ed', 'Abstract new deaths within 60 days of a positive test rate by death date', 'Rate per 100,000 people of the daily number of deaths of people who had a positive test for COVID-19 and either died within 60 days of their first positive test or have COVID-19 mentioned on their death certificate. Data are shown by the dates the deaths occurred.', '2021-09-03 13:19:52.942908', false),
        ('a458a6a5-7cca-4658-917f-9cd50a2e8cf5', 'Abstract new PCR tests change percentage by publish date', 'The percentage change in the number of new confirmed positive, negative or void polymerase chain reaction (PCR) test results during the latest 7-day period, as a percentage of the number for the previous, non-overlapping 7-day period. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 14:47:23.876479', false),
        ('c2d89368-d618-47aa-972b-94a2ee29d313', 'Abstract new cases LFD confirmed by PCR rolling sum by specimen date', 'The number of new cases (people who have had at least one positive COVID-19 test result) that were identified by a positive rapid lateral flow (LFD) test and confirmed by a positive polymerase chain reaction (PCR) test taken within 3 days, within rolling 7-day periods. Data are shown by the date the lateral flow test was taken.', '2021-09-02 15:25:30.711512', false),
        ('d8837194-8bdf-42ea-a8df-fbbab7674bf6', 'Abstract new deaths within 28 days of a positive test direction by publish date', 'The direction of the change in the number of deaths of people who had a positive test for COVID-19 and died within 28 days of their first positive test during the latest 7-day period compared to the previous, non-overlapping, 7-day period. Data are shown by the date the figures appeared in the published totals.

Positive changes mean numbers are increasing. These trends are shown with an upwards arrow. Negative changes mean numbers are decreasing. These trends are shown with a downwards arrow.', '2021-09-03 13:13:14.399743', false),
        ('e2ce773a-b79d-41e8-8c4b-022ec5ad8b73', 'Abstract new PCR tests direction by publish date', 'The direction of the change in the number of new confirmed positive, negative or void polymerase chain reaction (PCR) test results during the latest 7-day period compared to the previous, non-overlapping, 7-day period. Data are shown by the date the figures appeared in the published totals. 

Positive changes mean numbers are increasing. These trends are shown with an upwards arrow. Negative changes mean numbers are decreasing. These trends are shown with a downwards arrow.', '2021-09-03 14:49:16.330379', false),
        ('77be666f-addb-4085-a71e-f4d073076e88', 'Abstract new PCR tests rolling sum by publish date', 'The number of new confirmed positive, negative or void polymerase chain reaction (PCR) test results within rolling 7-day periods. Data are shown by the date the figures appeared in the published totals.', '2021-09-03 14:50:59.272669', false),
        ('124dcfde-3780-4f57-ac9f-841470f32177', 'Abstract new pillar 1 tests by publish date', 'Daily numbers of new confirmed positive, negative or void COVID-19 tests conducted under pillar 1 – NHS and PHE COVID-19 testing. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:54:58.514959', false),
        ('dc155672-9a11-473d-9af3-a76e798d6b02', 'New cases', 'Daily number of new people who have had at least one positive COVID-19 test result. People who test positive more than once are only counted once.', '2021-08-19 14:46:38.098581', false),
        ('00235aa2-393f-4bc1-a785-843795666898', 'Pillar 1 testing capacity', 'Projected lab capacity for NHS, PHE and Roche labs for England and lab capacity from devolved administrations.', '2021-08-19 15:14:38.911135', false),
        ('3ce985c7-e40f-439c-a8e9-3c253ff086c5', 'R number and growth rate (metadata)', 'The R number, or reproduction number, is the average number of secondary infections produced by a single infected person. If R is 2 then, on average, each infected person infects 2 more people. If R is 0.5 then, on average, for each 2 infected people there will be only 1 new infection. The R number indicates the direction of change, with a value greater than 1 meaning the epidemic is growing and a value less than 1 meaning the epidemic is shrinking. However, the R number alone does not indicate how quickly an epidemic is changing because it does not take into account the time taken for each transmission to take place.

The infection growth rate is an approximation of the percentage change in the number of infections each day and indicates how quickly the number of infections is changing. If the infection growth rate is greater than 0 (+ positive), then the number of new infections per day is increasing. If the infection growth rate is less than 0 (- negative) then the number of new infections per day is decreasing. The size of the infection growth rate indicates the speed of this change.

The R number and infection growth rate are available as estimated ranges for the UK, the four nations and NHS England Regions and represent the transmission of COVID-19 over the past few weeks due to the time delay between someone being infected, experiencing symptoms and needing healthcare. As they are averages over very different epidemiological situations they should be regarded as a guide to the general trend rather than a description of the epidemic state.

When the numbers of cases or deaths fall to low levels and/or there is a high degree of variability in transmission across a region, then care should be taken when interpreting estimates of R and the infection growth rate. For example, a significant amount of variability across a region due to a local outbreak may mean that a single average value does not accurately reflect the way infections are changing throughout that region.

A full definition of the R number and growth rate can be found on [GOV.UK](https://www.gov.uk/guidance/the-r-number-in-the-uk)', '2021-06-30 11:00:56.231202', false),
        ('a5ba9e44-defb-4159-9f39-dc34bb35edb5', 'Patients in hospital - metadata', '# Patients in hospital - metadata
The UK figure is the sum of the figures for England, Northern Ireland, Scotland and Wales. It can only be calculated when all nations'' data are available.
The definitions are not consistent between the 4 nations.
Data are not reported by each nation every day. This data is classified as management information rather than official statistics.
Use caution when interpreting the data as there are inconsistencies between the 4 nations:

-	definitions are not always consistent
-	for England, no revisions have been made to the dataset. When known errors come to light, Trusts make the appropriate correction in the next day’s data. For Wales, Northern Ireland and Scotland, historical data are subject to revision

## England
England data is the number of people currently in hospital with confirmed COVID-19 at 8am. Confirmed COVID-19 patients are patients who have tested positive for COVID-19 in the past 14 days through a polymerase chain reaction (PCR) test. 
Data include all hospitals and are reported daily by Trusts to NHS England and NHS Improvement.
Daily data at NHS Trust level is reported weekly on Thursdays.

## Northern Ireland
Northern Ireland data is the number of people currently in hospital with confirmed COVID-19 at midnight. 
Data are routinely uploaded to the Regional Data Warehouse from the Patient Administrative System (PAS), a patient-level administrative data source from health and social care hospitals in Northern Ireland. Data are sourced from the Regional Data Warehouse by the Department of Health Northern Ireland and published daily Monday to Friday. 

## Scotland
Scotland data is the number of people in hospital with confirmed COVID-19 at 8am the day before reporting.
Confirmed COVID-19 patients are patients who first tested positive in hospital or in the 14 days before admission. Patients stop being included after 28 days in hospital (or 28 days after first testing positive, if this is after admission). 
This figure includes all patients in hospitals, including in intensive care, community, mental health and long stay hospitals
Data are provided daily to the Scottish Government by NHS boards and are reported daily Monday to Friday. 

## Wales
Wales data includes:

-	patients in hospital with confirmed COVID-19
-	patients in hospital recovering from COVID-19

Confirmed COVID-19 patients are patients who have received a positive test result for COVID-19 in the past 14 days. Recovering patients are patients who were COVID-19 positive in hospital and who showed no symptoms for 14+ days but remained in hospital on a COVID-19 treatment pathway, often for rehabilitation.
Data include all acute hospitals. 
Data are collected in a daily return from Digital Health and Care Wales (DHCW) and are reported daily Monday to Friday by StatsWales.', '2021-08-19 15:00:38.608974', false),
        ('a9768ee6-6b22-41a2-b934-d13ea628b1af', '7-day rolling averages', '# 7-day rolling average

Each day''s figures are combined with the previous 3 days and the following 3 days. The average (mean) of all 7 days'' figures is shown.

If the most recent days'' data are incomplete, the final few points in the rolling average series are not shown, as the averages will increase when data are complete.', '2021-07-22 13:13:21.676768', false),
        ('7249f41c-1f19-47a0-a6a2-1c0e0354a712', 'Daily change in reported cases', 'Number of people with at least one positive COVID-19 virus test result broken down to show the number of these cases that had been previously reported and the latest daily change. The daily change includes newly reported cases, plus adjustments to previously reported cases which may be negative.', '2021-06-23 09:55:32.818860', false),
        ('d3373ca1-e15f-442a-b255-41f56aba18e7', 'Cumulative number of people tested by age and sex', 'Total number of people with at least one positive or negative COVID-19 virus test result since the start of the pandemic, by age and sex, and equivalent rates per 100,000 resident population. Some test results have missing age or sex, so the sum of the subgroups does not equal the total people tested for the area.', '2021-06-23 09:56:16.595894', false),
        ('80ac7006-df96-4220-bc3b-2c9223d8b18a', 'Daily patients in hospital', 'Daily patients in hospital.', '2021-07-19 10:27:27.326314', false),
        ('4c442d23-8548-4e43-ab32-6b9e928bb32f', 'Vaccination uptake (old)', 'Number of people and percentage of the population aged 18 and over who have received a COVID-19 vaccination.', '2021-08-10 12:44:35.213905', false),
        ('b4d9e739-04b2-4782-b522-d807dc1f6d14', 'Deaths ONS - intro', 'Deaths of people whose death certificate mentioned COVID-19 as one of the causes. 

Data can be presented by when someone died (date of death) or when the death was registered (date registered):

- deaths by date of death - each death is assigned to the date that the person died, however long it took for the death to be reported. Previously reported data are therefore continually updated
- deaths by registration date – deaths presented by the date the death was registered

Provisional counts of the number of deaths registered in England and Wales where COVID-19 is mentioned as a cause on the death certificate. The deceased may not have had a confirmed positive test for COVID-19. People who had COVID-19 but did not have it mentioned on their death certificate as a cause of death are excluded.', '2021-08-19 14:18:31.524964', false),
        ('6bacf27c-650a-49d5-b48e-65c85566fbdf', 'Abstract new pillar 4 tests by publish date', 'Daily numbers of new confirmed positive, negative or void COVID-19 tests conducted under pillar 4 – COVID-19 testing for national surveillance. Data are shown by the date the figures appeared in published totals.', '2021-09-03 14:53:14.905622', false),
        ('47e52551-3ca0-48ba-9917-3eccaf26e712', 'Rate calculations', '# Rate calculation
Rates are calculated in order to compare areas or population groups of different sizes. All rates currently presented on this website are crude rates expressed per 100,000 people, ie the count (eg cases or deaths) is divided by the denominator population and then multiplied by 100,000, without any adjustment for other factors.

Populations used are Office for National Statistics 2019 mid-year estimates, except for NHS Regions, for which 2019 estimates are not yet available, so 2018 mid-year estimates are used.', '2021-09-20 14:44:38.675054', false),
        ('a1103e2a-454e-49f4-9e44-a98c3cdfdccb', 'COVID-19 patients in mechanical ventilated beds', 'The UK figure is the sum of the four nations'' figures and can only be calculated when all nations'' data are available. Data are not reported by each nation every day. Caution is needed when interpreting the data as the definitions are not always consistent between the four nations.

# England
England figures are the numbers of patients in beds which are capable of delivering mechanical ventilation and includes Nightingale hospitals. Data are reported daily by trusts to NHS England and NHS Improvement. The data collected is classified as management information. It has been collected on a daily basis with a tight turn round time. No revisions have been made to the dataset, where known errors have come to light trusts have made the appropriate correction in the following day’s data. Any analysis of the data should be undertaken with this in mind.

Daily data at NHS Trust level is reported weekly on Thursdays, in line with NHS England reporting.

# Northern Ireland
Northern Ireland include suspected COVID-19 patients in their COVID-19 patient count for dates prior to 13 April 2020. Their figures are the numbers of patients in beds which are capable of delivering mechanical ventilation.

# Scotland
Scotland include suspected COVID-19 patients in their COVID-19 patient count. Their figures include people in intensive care units and may include a small number of patients who are not on mechanical ventilation. On 11 September 2020 the data were updated to exclude people (in larger NHS Boards) who had previously tested positive for COVID-19 but remain in intensive care units for another reason.

# Wales
Wales include suspected COVID-19 patients in their COVID-19 patient count. Their figures include invasive ventilated beds in a critical care setting, plus those outside of a critical care environment and also include surge capacity.

From 19 October 2020, specialist critical care beds have been included in these figures. From 13 November 2020, only critical care beds that could be staffed are included as available.

From 18 January 2021, no patients occupying an invasive ventilated bed (critical care bed) should be counted as “recovering” COVID-19 patients as they are still requiring a high level of care. Any patient previously reported as “recovering” will now be counted under “confirmed”. This will result in an increase in the number of invasive ventilated beds occupied by “confirmed” COVID-19 patients (an increase of around 14 patients at the point of implementation) and no invasive ventilated beds showing as occupied by “recovering” patients. This change will have no impact on the total number of COVID-19 related patients.', '2021-06-23 15:34:09.588942', false),
        ('964f1127-992f-4ce9-a8d5-c7a280e40320', 'Daily number of cases and 7-day rates by age group (0-59 and 60+ years)', 'Daily number of people with at least one positive COVID-19 virus test result, by specimen date, and seven day rates of cases per 100,000 population, grouped by age 0-59 and 60+ years.', '2021-06-23 09:38:48.656456', false),
        ('a97d2348-66b1-4dac-8abb-71eded0ab2f2', 'People tested', 'People tested includes all individuals who have had one or more lab-confirmed positive or negative COVID-19 PCR test results.

COVID-19 cases are identified by taking specimens from people and sending these specimens to laboratories around the UK for PCR swab testing. If the test is positive, this is a referred to as a lab-confirmed case. If a person has had more than one test result they are only counted as one person tested. If any of their tests were positive they count as a case. If all their tests were negative they count as one person tested negative. The people tested figure is the sum of cases and people tested negative.

Initially only pillar 1 tests were included but pillar 2 tests have been included since 15th June for Scotland, 26th June for Northern Ireland, 2nd July for England and 14th July for Wales (see below).

Details of the processes for counting cases in the devolved administrations are available on their websites:

- [Scottish Government coronavirus information](https://www.gov.scot/coronavirus-covid-19/)
- [Public Health Wales coronavirus information](https://public.tableau.com/profile/public.health.wales.health.protection#!/vizhome/RapidCOVID-19virology-Public/Headlinesummary)
- [Northern Ireland Department of Health coronavirus information](https://app.powerbi.com/view?r=eyJrIjoiZGYxNjYzNmUtOTlmZS00ODAxLWE1YTEtMjA0NjZhMzlmN2JmIiwidCI6IjljOWEzMGRlLWQ4ZDctNGFhNC05NjAwLTRiZTc2MjVmZjZjNSIsImMiOjh9)', '2021-06-30 10:56:23.363150', false),
        ('ee1389d8-e514-4668-9f98-7e26311de6cf', 'Tests processed', 'There are three types of test being carried out:

- Lab-based virus tests that test for the presence of SARS-CoV-2 virus. These include lab-based pillar 1 and 2 tests and virus tests undertaken in pillar 4.
- Antibody serology tests that test for the presence of COVID-19 antibodies. These include pillar 3 tests and antibody serology tests undertaken in pillar 4. All pillar 4 antibody testing is included in reported UK totals only.
- Rapid lateral flow virus tests that test for the presence of SARS-CoV-2 virus. These are swab tests that give results in less than an hour, without needing to go to a laboratory. Rapid lateral flow test data are currently available for England only.

The number of tests conducted counts all tests that have remained within the control of the programme and those that have been sent out and subsequently returned to be processed in a lab. These include tests which are negative and positive and may also include tests which were inconclusive (void). Reasons for void tests include:

-	The testing within the laboratory was unable to give an exact answer as the test gave an equivocal result (it wasn’t clearly negative or positive).
-	There was an issue with the sample when it is delivered to the laboratory making it impossible to return a result, such as the tube was damaged or leaked, the test could not be scanned, the instructions were not followed or the swab was not put in the tube. Due to the issue with the sample, some of these tests may not have been processed by the lab at all.

Tests are counted at the time at which they were processed.

# UK
UK data are available for virus testing and antibody serology testing.

# England and Wales
The number of tests conducted by test result in England is published as part of the weekly [NHS Test and Trace Statistics](https://www.gov.uk/government/collections/nhs-test-and-trace-statistics-england-weekly-reports).
This data contains rapid lateral flow tests reported through the existing National Testing Programme digital infrastructure and doesn''t include tests conducted where the tests were not registered digitally. In future, all rapid lateral flow tests will be reported via the existing National Testing Programme digital infrastructure and will be included.', '2021-07-21 10:49:39.995002', false),
        ('6b243be5-6a81-4a6f-9cc3-83ac8640d4f0', '7-day totals', '# 7-day total

The total number of events (cases, deaths, tests, hospitalisations, vaccinations) reported in the latest 7-day period for which data are complete. 

The 7-day totals use the event date (for example, the date a person took a test for COVID-19), not the reporting date (for example, the date a positive test result was reported). This means the most recent 5 days'' data are considered incomplete. The latest 7-day period will be 5 days behind other website data updates.

## Change in 7-day total

The difference between the 7-day totals for the latest period and the 7-day total for the previous, non-overlapping, 7-day period. 
The percentage change in the 7-day total shows this change as a percentage of the previous 7-day total. 

Positive changes mean numbers are increasing. These trends are shown with an upwards arrow. Negative changes mean numbers are decreasing. These trends are shown with a downwards arrow.

Some trends might be seen as good (for example, falling cases) or bad (for example, growing cases).  These 7-day change figures are shown on a green background if the metric is moving in the desired direction or a red background if the metric is moving in an undesirable direction. If a trend is not obviously good or bad (for example, changes in numbers of tests conducted), the 7-day change figures are shown on a grey background. 


If an area''s previous 7-day count was suppressed (counts of 0-2) but the latest 7-day count is displayed, a value of 2 is used for the previous 7-day count when calculating the change in the 7-day count. 

Where the case rate is shown on a scale compared to an average (this comparison is shown at MSOA level only), the average used is the median of all areas.', '2021-07-22 13:32:08.782371', false);