// import { MigrationInterface, QueryRunner, Table } from 'typeorm'

// export class CreateSubjectsMigration implements MigrationInterface {
//   name = this.constructor.name.concat('1657896596925')

//   public async up(queryRunner: QueryRunner): Promise<void> {
//     await queryRunner.createTable(
//       new Table({
//         name: 'subjects',
//         columns: [
//           { name: 'id', type: 'uuid', isPrimary: true },
//           { name: 'name', type: 'varchar', isNullable: false },
//           { name: 'created_at', type: 'timestamp', default: 'now()' }
//         ]
//       }),
//       true
//     )
//   }

//   public async down(queryRunner: QueryRunner): Promise<void> {
//     await queryRunner.dropTable('subjects', true)
//   }
// }
